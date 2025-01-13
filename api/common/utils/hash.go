package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(password), 1)

	return string(res), err
}

func ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
