// Package payment REST API methods exposed by merchant endpoint

package payment

import (
	"context"
	bc "github.com/eldadj/dgpg/controllers"
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/eldadj/dgpg/internal/merchant"
)

type BaseController struct {
	// set in authorize call and used by other methods in the package.
	AuthorizedId string
	bc.BaseController
}

// SetMerchantId sets the request merchant id from authenticated merchant's id in context
func SetMerchantId(ctx context.Context, req *payment.Request) error {
	var ok bool
	// return error if not matched
	if req.MerchantId, ok = merchant.FromContext(ctx); !ok {
		return errors.ErrMerchantLoad
	}

	return nil
}
