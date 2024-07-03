package models

import (
	"errors"
)

func passwordMatch(password string, confirmPassword string) error {
	if password != confirmPassword {
		return errors.New("passwords do not match")
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	return nil
}

func alreadyRegistered(username string) error {
	db, err := Connection()
	if err != nil {
		return err
	}
	var count int
	err = db.QueryRow("SELECT COUNT(username) FROM users WHERE username=?", username).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("username already registered")
	}
	return nil
}

func Registration(username string, password string, confirmPassword string) error {
	err := alreadyRegistered(username)
	if err!=nil {
		return err
	}
	err = passwordMatch(password, confirmPassword)
	if err!=nil {
		return err
	}
	err = validatePassword(password)
	if err!=nil {
		return err
	}
	hashedPassword, err := HashPass(password)
	if err!=nil {
		return err
	}
	
	db, err := Connection()
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO users (username, password, isadmin) VALUES (?, ?, 0)", username, hashedPassword)
	db.Close()
	if err != nil {
		return err
	}
	return nil
}