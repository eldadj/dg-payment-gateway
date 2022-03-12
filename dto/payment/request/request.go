// Package capture request/response objects for capture endpoint
package request

import (
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/internal/errors"
)

type Request struct {
	payment.AuthorizeCode
	Amount float64 `json:"amount"`
}

func (r *Request) Validate() error {
	if r.Code == "" {
		return errors.ErrAuthorizeCodeInvalid
	}
	return nil
}
