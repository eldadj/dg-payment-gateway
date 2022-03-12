package credit_card

import (
	"database/sql"
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/internal/errors"
	"gorm.io/gorm"
)

func AmountCurrency(tx *gorm.DB, creditCardId int64) (payAC payment.AmountCurrency, err error) {
	//now get current amount/currency for the card
	var amount sql.NullFloat64
	var currency sql.NullString
	row := tx.Raw(`select currency_code, current_amount from credit_card where credit_card_id = ?`, creditCardId).Row()
	err = row.Scan(&currency, &amount)
	if !amount.Valid || !currency.Valid {
		return payAC, errors.ErrCreditCardLoadAmountCurrency
	}
	if err != nil {
		err = errors.LogError(err, errors.ErrCreditCardLoadAmountCurrency)
	}
	payAC.Currency = currency.String
	payAC.Amount = amount.Float64

	return
}
