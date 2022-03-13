package payment

import (
	"github.com/theplant/luhn"
	"strconv"
	"strings"
)

//IsCreditCardLuhnValid checks if the card passes luhn test
func IsCreditCardLuhnValid(cardNo string) bool {
	cNo := strings.ReplaceAll(cardNo, " ", "")
	cardNumber, err := strconv.Atoi(cNo)
	return err == nil && luhn.Valid(cardNumber)
}
