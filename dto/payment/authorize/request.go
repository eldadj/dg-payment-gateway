// Package authorize request/response objects for authorize endpoint
package authorize

import "github.com/eldadj/dgpg/dto/payment"

type Request struct {
	//MerchantId int64 `json:"-"`
	payment.CreditCard
	payment.AmountCurrency
}

func (r *Request) Validate() error {
	if err := r.CreditCard.Validate(); err != nil {
		return err
	}
	if err := r.AmountCurrency.Validate(); err != nil {
		return err
	}
	return nil
}
