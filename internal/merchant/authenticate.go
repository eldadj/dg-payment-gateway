package merchant

import (
	"golang.org/x/crypto/bcrypt"
)

func IsPasswordHashValid(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
