package credit_card

import (
	"database/sql"
	"errors"
	errors2 "github.com/eldadj/dgpg/internal/errors"
	"github.com/eldadj/dgpg/models"
	"gorm.io/gorm"
)

func GetCardNumber(creditCardId int64) (cardNo string, err error) {
	err = models.ExecDBFuncReadOnly(func(tx *gorm.DB) error {
		var err error
		var cardNoNull sql.NullString
		row := tx.Raw(`select card_no from credit_card where credit_card_id = ?`, creditCardId).Row()
		err = row.Scan(&cardNoNull)
		if cardNoNull.Valid && (err == nil || errors.Is(err, gorm.ErrRecordNotFound)) {
			cardNo = cardNoNull.String
			return nil
		}
		return errors2.LogError(err, errors2.ErrCreditCardLoad)
	})
	return
}
