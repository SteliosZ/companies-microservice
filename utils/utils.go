package utils

import (
	"net/mail"
	"os"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt/v5"
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

func ComparePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return err
	}

	return nil
}

// GenerateJWT generates a jwt token
func GenerateJWT(id *uuid.UUID) (string, error) {
	// Create claims
	claims := jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	// Generate Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with given secret from .env
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return t, nil
}
