package capture

import (
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/dto/payment/request"
	"github.com/eldadj/dgpg/dto/payment/response"
	"github.com/eldadj/dgpg/internal"
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/eldadj/dgpg/models"
	"github.com/eldadj/dgpg/models/authorize"
	"github.com/eldadj/dgpg/models/credit_card"
	"gorm.io/gorm"
)

type capture struct {
	CaptureId   int64 `gorm:"primaryKey"`
	Amount      float64
	AuthorizeId int64
}

func (*capture) TableName() string {
	return "capture"
}

//DoCapture handles the capture process
func DoCapture(req request.Request) (resp response.Response, err error) {
	var auth authorize.Authorize

	// 1 get authorize record
	if auth, err = authorize.Get(req.Code); err != nil {
		return resp, err
	}

	// get credit card and do checks on the number
	cardNo, err := credit_card.GetCardNumber(auth.CreditCardId)
	if err != nil {
		return resp, err
	}
	if err = internal.CardNoCheck(cardNo, internal.CaptureFailCard, errors.ErrCaptureCreditCardFailed); err != nil {
		return resp, err
	}

	//2 check authorize record can still be captured
	if !auth.CanCapture() {
		return resp, errors.ErrAuthorizeCannotCapture
	}

	//3 check amount to capture won't exceed authorized amount
	// done in a new readonly transaction in case another capture comes before we are done
	err = models.ExecDBFuncReadOnly(func(tx *gorm.DB) error {
		totalAmountCaptured, err := models.TotalAmountCaptured(tx, auth.AuthorizeId)
		if totalAmountCaptured+req.Amount > auth.Amount {
			return errors.ErrCaptureAmountExceedsAuthorizeAmount
		}
		return err
	})
	if err != nil {
		return resp, err
	}

	err = models.ExecDBFunc(func(tx *gorm.DB) error {
		var err error
		//4 save capture data
		if err = Add(tx, auth.AuthorizeId, req.Amount); err != nil {
			return err
		}

		//5 update authorize record status
		if err = auth.UpdateStatusCapture(tx); err != nil {
			return err
		}
		//6 update credit card amount and current value
		var payAC payment.AmountCurrency
		if payAC, err = credit_card.UpdateAmount(tx, auth.CreditCardId, req.Amount); err != nil {
			return err
		}
		resp.Success = true
		resp.AmountCurrency = payAC

		return err
	})

	return
}
