package authorize

import (
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/eldadj/dgpg/models"
	"gorm.io/gorm"
)

func Get(authorizeCode string, merchantId int64) (a Authorize, err error) {
	err = models.ExecDBFuncReadOnly(func(tx *gorm.DB) error {
		result := tx.Find(&a, "authorize_code = ?", authorizeCode)
		if result.RowsAffected == 0 {
			return errors.ErrAuthorizeCodeNotFound
		}
		// check if for the merchant
		if a.MerchantId != merchantId {
			return errors.ErrInvalidMerchantAuthorizeCode
		}
		return nil
	})

	return
}
