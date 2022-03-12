package capture

import (
	"github.com/eldadj/dgpg/internal/errors"
	"gorm.io/gorm"
)

func Add(tx *gorm.DB, authorizeId int64, amount float64) (err error) {
	if amount <= 0 {
		return errors.ErrCaptureAmount
	}
	c := capture{
		Amount:      amount,
		AuthorizeId: authorizeId,
	}
	result := tx.Create(&c)
	if result.Error != nil {
		return errors.LogError(result.Error, errors.ErrCapture)
	}

	return
}
