package void

import (
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/internal/errors"
)

type Request struct {
	payment.AuthorizeCode
	// so we have access to MerchantId
	payment.Request
}

func (r *Request) Validate() error {
	if r.Code == "" {
		return errors.ErrAuthorizeCodeInvalid
	}
	return nil
}
