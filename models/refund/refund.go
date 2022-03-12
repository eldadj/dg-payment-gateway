// Package refund
package refund

import (
	"database/sql"
	"github.com/eldadj/dgpg/dto/payment/request"
	"github.com/eldadj/dgpg/dto/payment/response"
	"github.com/eldadj/dgpg/internal"
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/eldadj/dgpg/models"
	"github.com/eldadj/dgpg/models/authorize"
	"github.com/eldadj/dgpg/models/credit_card"
	"gorm.io/gorm"
)

func DoRefund(req request.Request) (resp response.Response, err error) {
	var auth authorize.Authorize

	// 1 get authorize record
	if auth, err = authorize.Get(req.Code, req.MerchantId); err != nil {
		return resp, err
	}

	// get credit card and do checks on the number
	cardNo, err := credit_card.GetCardNumber(auth.CreditCardId)
	if err != nil {
		return resp, err
	}
	if err = internal.CardNoCheck(cardNo, internal.RefundFailCard, errors.ErrRefundCreditCardFailed); err != nil {
		return resp, err
	}

	//2 check authorize record can still be refunded
	if !auth.CanRefund() {
		return resp, errors.ErrAuthorizeCannotRefund
	}

	//3
	err = models.ExecDBFunc(func(tx *gorm.DB) error {
		var err error
		//get the last capture record that hasn't been voided and has same amount
		//this ensures we refund most recent capture of same amount
		var captureId int64
		row := tx.Raw(`select capture_id from capture where authorize_id = ? and amount = ? and refunded = false 
			order by capture_id desc limit 1`, auth.AuthorizeId, req.Amount).Row()
		if err = row.Scan(&captureId); err != nil {
			return errors.LogError(err, errors.ErrRefundAmount)
		}
		//set refund
		if err = tx.Exec(`update capture set refunded = true where capture_id = ?`, captureId).Error; err != nil {
			return errors.LogError(err, errors.ErrRefund)
		}
		//TODO: collapse next 2 execs into a single call
		//set we have a refund
		if err = tx.Exec(`update authorize set has_refund = true where authorize_id = ? and has_refund = false`,
			auth.AuthorizeId).Error; err != nil {
			return errors.LogError(err, errors.ErrRefund)
		}
		//set status = r if we have no more pending un-refunded captures
		var unRefundCount sql.NullInt64
		row = tx.Raw(`select count(*) refund_count from capture where refunded = false and authorize_id = ?`,
			auth.AuthorizeId).Row()
		err = row.Scan(&unRefundCount)
		if unRefundCount.Valid && unRefundCount.Int64 > 0 {
			if err = tx.Exec(`update authorize set status = 'R' where authorize_id = ?`, auth.AuthorizeId).Error; err != nil {
				return errors.LogError(err, errors.ErrRefund)
			}
		}

		//create a refund record
		//TODO: does this need to be as an orm? its simple enough
		err = tx.Exec(`insert into refund(authorize_id)values(?)`, auth.AuthorizeId).Error

		//update credit card amount
		payAC, err := credit_card.UpdateAmount(tx, auth.CreditCardId, -req.Amount)
		if err != nil {
			return err
		}
		resp.Amount = payAC.Amount
		resp.Currency = payAC.Currency

		return err
	})

	return
}
