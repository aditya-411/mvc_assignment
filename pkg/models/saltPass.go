package models

import (
	"os"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

// HashPass salts the given password using bcrypt
func HashPass(password string) (string, error) {
	// Generate a salt for bcrypt
	saltRounds := os.Getenv("salt_rounds")
	saltRoundsInt, err := strconv.Atoi(saltRounds)
	if err != nil {
		return "", err
	}
	salt, err := bcrypt.GenerateFromPassword([]byte(password), saltRoundsInt)

	if err!=nil{
		return "", err
	} 
	hashedPassword := string(salt)
	return hashedPassword, nil
}