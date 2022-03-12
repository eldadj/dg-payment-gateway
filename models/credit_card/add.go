package credit_card

import (
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/internal/errors"
	"gorm.io/gorm"
)

func Add(tx *gorm.DB, payCC payment.CreditCard, payAC payment.AmountCurrency) (
	creditCardId int64, ac payment.AmountCurrency, err error) {
	var cc creditCard
	//card is tied to authorize via credit_card_id
	// we add it to db
	cc = creditCard{
		OwnerName:     payCC.OwnerName,
		CardNo:        payCC.Number,
		ExpMonth:      payCC.ExpMonth,
		ExpYear:       payCC.ExpYear,
		CVV:           payCC.CVV,
		CurrencyCode:  payAC.Currency,
		CurrentAmount: payAC.Amount,
	}
	if err = tx.Create(&cc).Error; err != nil {
		err = errors.LogError(err, errors.ErrCreditCardSave)
	}

	creditCardId = cc.CreditCardId
	//get existing values
	ac.Amount = cc.CurrentAmount
	ac.Currency = cc.CurrencyCode

	return
}
