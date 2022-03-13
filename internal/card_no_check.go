package internal

import (
	"github.com/eldadj/dgpg/internal/errors"
	"strings"
)

// CardNoCheck do checks on card no
func CardNoCheck(cardNo string, againstCardNo string, returnErr error) (err error) {
	if cardNo == "" || againstCardNo == "" {
		if returnErr != nil {
			return returnErr
		}
		return errors.ErrCreditCardNoInvalid
	}
	//remove all spaces to make check easy
	cardNumber := strings.ReplaceAll(cardNo, " ", "")
	if strings.EqualFold(cardNumber, strings.ReplaceAll(againstCardNo, " ", "")) {
		return returnErr
	}
	return nil
}
