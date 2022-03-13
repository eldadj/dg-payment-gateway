package merchant

import (
	"golang.org/x/crypto/bcrypt"
)

//IsPasswordHashValid validates a password using its hash
func IsPasswordHashValid(password, hash string) bool {
	//be safe, invalid if any is blank
	if password == "" || hash == "" {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
