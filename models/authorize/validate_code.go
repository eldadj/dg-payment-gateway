package authorize

import (
	"database/sql"
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/eldadj/dgpg/models"
	"gorm.io/gorm"
	"strings"
)

// ValidateAuthorizeCode returns nil if the code exists and hasn't been finalized
func ValidateAuthorizeCode(authorizeCode string) error {
	err := models.ExecDBFunc(func(tx *gorm.DB) error {
		row := tx.Raw(`select authorize_id, status from authorize where authorize_code = ?`,
			authorizeCode).Row()
		var authorizeId sql.NullInt64
		var status sql.NullString
		row.Scan(&authorizeId, &status)
		if !authorizeId.Valid || !status.Valid {
			return errors.ErrAuthorizeCodeNotFound
		}
		// code is still processable if new or in process
		if !(strings.EqualFold("N", status.String) || strings.EqualFold("P", status.String)) {
			return errors.ErrAuthorizeCodeInvalidStatus(status.String)
		}
		return nil
	})
	return err
}
