package payment

import "errors"

var ErrAmountCurrencyInvalid = errors.New("amount/currency is invalid")

type AmountCurrency struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

func (a *AmountCurrency) Validate() error {
	if a.Amount <= 0 || a.Currency == "" {
		return ErrAmountCurrencyInvalid
	}
	return nil
}
