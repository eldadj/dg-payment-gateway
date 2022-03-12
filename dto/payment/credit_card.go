package payment

import "errors"

var ErrCreditCardInvalid = errors.New("invalid credit card details")

type CreditCard struct {
	OwnerName string `json:"owner_name"`
	Number    string `json:"number"`
	ExpMonth  int    `json:"exp_month"`
	ExpYear   int    `json:"exp_year"`
	CVV       string `json:"cvv"`
}

func (c *CreditCard) Validate() error {
	if c.OwnerName == "" || c.Number == "" || c.ExpMonth <= 0 || c.ExpYear <= 0 || c.CVV == "" {
		return ErrCreditCardInvalid
	}
	return nil
}
