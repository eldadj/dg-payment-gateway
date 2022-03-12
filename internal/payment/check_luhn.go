package payment

import (
	"github.com/theplant/luhn"
	"strconv"
)

func IsCreditCardLuhnValid(cardNo string) bool {
	return true
	cardNumber, err := strconv.Atoi(cardNo)
	return err == nil && luhn.Valid(cardNumber)
}
