package models

import (
	"database/sql"
	"gorm.io/gorm"
)

//TotalAmountCaptured get total amount already captured and not refunded
// return 0 when not found
func TotalAmountCaptured(tx *gorm.DB, authorizeId int64) (amount float64, err error) {
	var amountNull sql.NullFloat64
	row := tx.Raw(`select sum(amount) total_amount from capture where authorize_id = ? and refunded = false`,
		authorizeId).Row()
	err = row.Scan(&amountNull)
	if amountNull.Valid {
		return amountNull.Float64, err
	}
	return 0, err
}
