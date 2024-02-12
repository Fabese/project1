package db

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	if err != nil {

	}
	return string(bytes), nil
}
