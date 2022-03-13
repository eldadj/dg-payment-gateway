package payment

import "github.com/eldadj/dgpg/internal/errors"

type CreditCard struct {
	OwnerName string `json:"owner_name"`
	Number    string `json:"number"`
	ExpMonth  int    `json:"exp_month"`
	ExpYear   int    `json:"exp_year"`
	CVV       string `json:"cvv"`
}

func (c *CreditCard) Validate() error {
	if c.OwnerName == "" || c.Number == "" || c.ExpMonth <= 0 || c.ExpYear <= 0 || c.CVV == "" {
		return errors.ErrCreditCardInvalid
	}
	return nil
}
