package authorize

import (
	"database/sql"
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/eldadj/dgpg/models"
	"github.com/eldadj/dgpg/models/credit_card"
	"gorm.io/gorm"
)

func CreditCardAmountCurrency(authorizeCode string) (payAC payment.AmountCurrency, err error) {
	err = models.ExecDBFuncReadOnly(func(tx *gorm.DB) error {
		var err error
		//get card id
		var ccId sql.NullInt64
		row := tx.Raw(`select credit_card_id from authorize where authorize_code = ?`, authorizeCode).Row()
		row.Scan(&ccId)
		if !ccId.Valid {
			return errors.ErrAuthorizeCodeNotFound
		}
		//get the data
		payAC, err = credit_card.AmountCurrency(tx, ccId.Int64)

		return err
	})
	return
}
