package credit_card

import (
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/internal/errors"
	"gorm.io/gorm"
)

// UpdateAmount reduces the current credit_card amount by amount
func UpdateAmount(tx *gorm.DB, creditCardId int64, amount float64) (payAC payment.AmountCurrency, err error) {
	result := tx.Model(creditCard{}).
		Where("credit_card_id = ?", creditCardId).
		UpdateColumn("current_amount", gorm.Expr("current_amount - ?", amount))
	if result.RowsAffected != 1 {
		return payAC, errors.LogError(result.Error, errors.ErrCreditCardUpdateAmount)
	}
	var cc *creditCard
	result = tx.Find(&cc, "credit_card_id = ?", creditCardId)
	if result.RowsAffected != 1 {
		return payAC, errors.LogError(result.Error, errors.ErrCreditCardLoadAmountCurrency)
	}
	payAC.Amount = cc.CurrentAmount
	payAC.Currency = cc.CurrencyCode

	return
}
