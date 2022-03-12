package authorize

import (
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/eldadj/dgpg/models"
	"gorm.io/gorm"
)

func Get(authorizeCode string) (a Authorize, err error) {
	err = models.ExecDBFuncReadOnly(func(tx *gorm.DB) error {
		result := tx.Find(&a, "authorize_code = ?", authorizeCode)
		if result.RowsAffected == 0 {
			return errors.ErrAuthorizeCodeNotFound
		}
		return nil
	})

	return
}
