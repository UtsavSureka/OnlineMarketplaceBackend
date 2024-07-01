package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", errors.New("cannot convert to hashed password")
	}
	return string(bytes), err

}

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, errors.New("incorrect Password")
	}
	return true, nil
}
