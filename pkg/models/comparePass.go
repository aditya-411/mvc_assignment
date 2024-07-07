package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func ComparePass(username string, password string) (bool, error) {
	db, err := Connection()
	if err != nil {
		return false, fmt.Errorf("error connecting to the database: %s", err)
	}
	defer db.Close()

	selectSql := "SELECT password FROM users WHERE username=?"
	rows, err := db.Query(selectSql, username)

	if err != nil {
		return false, fmt.Errorf("error querying the database: %s", err)
	}

	var password_fetched string

	if rows.Next() {
		err = rows.Scan(&password_fetched)
		if err != nil {
			return false, fmt.Errorf("error scanning the result set: %s", err)
		}
	} else {
		return false, fmt.Errorf("username not found in the database")
	}


	if bcrypt.CompareHashAndPassword([]byte(password_fetched), []byte(password)) != nil {
		return false, fmt.Errorf("passwords do not match")
	}

	return true , nil
}