package void

import (
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/internal/errors"
)

type Request struct {
	payment.AuthorizeCode
}

func (r *Request) Validate() error {
	if r.Code == "" {
		return errors.ErrAuthorizeCodeInvalid
	}
	return nil
}
