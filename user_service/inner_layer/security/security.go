package security

import "golang.org/x/crypto/bcrypt"

func IsFineCheckPasswordAndHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func GeneratePasswordHash(password string) (hash []byte, err error) {
	hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	return
}
