package utils

import (
	"golang.org/x/crypto/bcrypt"
)
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password string, requestPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(requestPassword))
    if err != nil {
        return err
    }
    return nil
}