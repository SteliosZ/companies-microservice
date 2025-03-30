package utils

import (
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

// isEmail functions as a type checker for email string validation
func CheckEmailValidity(email string) error {
	_, err := mail.ParseAddress(email)

	if err != nil {
		return err
	}
	return nil
}

func GenerateHash(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
