package authorize

import (
	"database/sql"
	"errors"
	mutil "github.com/eldadj/dgpg/internal/errors"
	"gorm.io/gorm"
)

// TotalAuthorizeAmounts determine max amount that can be authorized by checking pending authorizes and
func TotalAuthorizeAmounts(tx *gorm.DB, creditCardId int64) (float64, error) {
	//we assume card is authorise for a single merchant
	//TODO: refactor to also include the merchant_id
	row := tx.Raw(`select sum(amount) amount from authorize where upper(status) in ('N', 'P') 
		and credit_card_id = ?`, creditCardId).Row()
	var authorizedAmount sql.NullFloat64
	err := row.Scan(&authorizedAmount)
	if authorizedAmount.Valid || err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return authorizedAmount.Float64, nil
	}
	return 0, mutil.LogError(err, mutil.ErrAuthorizeAmountsAuthorize)
}
