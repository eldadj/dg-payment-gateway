package authorize

import (
	"context"
	"fmt"
	"github.com/eldadj/dgpg/dto"
	"github.com/eldadj/dgpg/dto/payment"
	auth "github.com/eldadj/dgpg/dto/payment/authorize"
	"github.com/eldadj/dgpg/internal"
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/eldadj/dgpg/internal/merchant"
	payment2 "github.com/eldadj/dgpg/internal/payment"
	"github.com/eldadj/dgpg/models"
	"strings"

	"gorm.io/gorm"
	"time"
)

type Authorize struct {
	AuthorizeId  int64 `gorm:"primaryKey"`
	MerchantId   int64
	CreditCardId int64
	Currency     string
	Amount       float64
	DateIn       time.Time `gorm:"default:null"` // set on db insert
	//TimeIn        time.Time `gorm:"default:null"` // set on db insert
	AuthorizeCode string
	//Refunded      bool `gorm:"default:false"`
	//Captured      bool `gorm:"default:false"`
	//n = new just created,
	//v = voided,
	//r = refunded,
	//p = capturing or when at least when capture has taken place
	//c = fully captured
	Status    string `gorm:"default:'N'"`
	HasRefund bool   `gorm:"default:false"`
}

func (*Authorize) TableName() string {
	return "authorize"
}

// BeforeCreate set the authorize_code when creating if not set
func (a *Authorize) BeforeCreate(tx *gorm.DB) (err error) {
	if a.AuthorizeCode == "" {
		//TODO: switch to nanoid
		a.AuthorizeCode = fmt.Sprintf("%d", time.Now().UnixNano())
	}
	return
}

// DoAuthorize handles creating authorise record and returning a new authorize_code
func DoAuthorize(ctx context.Context, req auth.Request) (resp auth.Response, err error) {
	//do some checks
	if err = internal.CardNoCheck(req.Number, internal.AuthorizeFailCard, errors.ErrAuthorizeCreditCardFailed); err != nil {
		return resp, err
	}
	//check card luhn
	cardNo := strings.ReplaceAll(req.Number, " ", "")
	if !payment2.IsCreditCardLuhnValid(cardNo) {
		return resp, errors.ErrCreditCardNoInvalid
	}

	err = models.ExecDBFunc(func(tx *gorm.DB) error {

		// get/create creditCard
		/*creditCardId, ccAmountCurrency, err := credit_card.Add(tx, req.CreditCard, req.AmountCurrency)
		if err != nil {
			return err
		}*/

		//TODO: verify that currency equals saved card currency

		//get all authorized amounts
		/*totalAuthorizeAmounts, err := TotalAuthorizeAmounts(tx, creditCardId)
		if err != nil {
			return err
		}
		cardBalance := ccAmountCurrency.Amount - totalAuthorizeAmounts

		//we ensure amount left in card can handle requested amount
		if cardBalance < req.Amount {
			return errors.ErrAuthorizeInsufficientCreditCardAmount
		}*/

		//get merchantid from context
		merchantId, ok := merchant.FromContext(ctx)
		if !ok {
			return errors.ErrAuthorizeFailed
		}

		//we store requested amount
		//ccAmountCurrency.Amount = req.Amount

		// add authorize record
		authorizeCode, err := Add(tx, merchantId, req)
		if err != nil {
			return err
		}
		resp = auth.Response{
			AuthorizeCode: payment.AuthorizeCode{
				Code: authorizeCode,
			},
			AmountCurrency: payment.AmountCurrency{
				Amount: req.Amount, Currency: req.Currency,
			},
			Response: dto.Response{
				Success: true,
			},
		}
		return nil
	})
	return
}

// UpdateStatusCapture set the status to 'C':fully captured if totalCapturedAmount == authorizeAmount
// else 'P': in progress (we still have capturable amount)
func (a *Authorize) UpdateStatusCapture(tx *gorm.DB) error {
	//get total amount already captured
	totalCaptured, _ := models.TotalAmountCaptured(tx, a.AuthorizeId)
	//totalCaptured := 0.0
	status := "P"
	if totalCaptured >= a.Amount {
		status = "C"
	}
	tx.Model(a).Update("status", status)
	return nil
}
