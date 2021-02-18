package crypto

import (
	"github.com/coposaja/bookstore-api/src/utils/rerr"
	"golang.org/x/crypto/bcrypt"
)

// HashAndSalt return hashed string
func HashAndSalt(password string) (string, rerr.RestError) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", rerr.NewInternalServerError(err.Error())
	}

	return string(hash), nil
}
