package authorize

import (
	"github.com/eldadj/dgpg/models"
	"gorm.io/gorm"
)

func CanCaptureAmount(authorizeCode string, amount float64) (canCaptureAmount bool, err error) {
	err = models.ExecDBFuncReadOnly(func(tx *gorm.DB) error {
		var err error
		//var captureAmount float64
		//row := tx.Raw(``)

		return err
	})

	return
}
