package models

import (
	// "fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// CreateJWT creates a new JWT token with the given claims and secret key
func GenerateJWT(username string) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}

	var isadmin bool
	err = db.QueryRow("SELECT isadmin FROM users WHERE username=?", username).Scan(&isadmin)
	if err != nil {
		return "", err
	}
	defer db.Close()
	
	claims := jwt.MapClaims{
		"username": username,
		"isadmin":  isadmin,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("jwt_secret")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}