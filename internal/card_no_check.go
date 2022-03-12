package internal

import "strings"

// CardNoCheck do checks on card no
func CardNoCheck(cardNo string, againstCardNo string, returnErr error) (err error) {
	//remove all spaces to make check easy
	cardNumber := strings.ReplaceAll(cardNo, " ", "")
	if strings.EqualFold(cardNumber, strings.ReplaceAll(againstCardNo, " ", "")) {
		return returnErr
	}
	return nil
}
