package void

import (
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/dto/payment/void"
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/eldadj/dgpg/models"
	"github.com/eldadj/dgpg/models/authorize"
	"gorm.io/gorm"
)

func Void(req void.Request) (resp void.Response, err error) {
	//var a authorize.Authorize
	a, err := authorize.Get(req.Code, req.MerchantId)
	if err != nil {
		return resp, err
	}
	if !a.CanVoid() {
		return resp, errors.ErrAuthorizeCannotVoid
	}

	//now void
	err = models.ExecDBFunc(func(tx *gorm.DB) error {
		var err error
		if result := tx.Exec(`update authorize set status = 'V' where authorize_code = ?`, req.Code); result.RowsAffected != 1 {
			return errors.LogError(result.Error, errors.ErrAuthorizeVoidFailed)
		}
		//get associated card amount/currency
		var payAC payment.AmountCurrency
		if payAC, err = authorize.CreditCardAmountCurrency(req.Code); err != nil {
			return err
		}
		resp.Success = true
		resp.AmountCurrency = payAC

		return err
	})

	return
}
