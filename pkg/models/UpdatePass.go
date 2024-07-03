package models

import (
	"fmt"
)

func UpdatePass(currentPass string, newPass string, confirmPass string, username string) string {
	correctCurrentPass, _ := ComparePass(username, currentPass)
	if !correctCurrentPass{
		return "Current password is incorrect"
	}
	if newPass != confirmPass {
		return "Passwords do not match"
	}
	if newPass == currentPass{
		return "New password cannot be the same as the current password"
	}
	err := validatePassword(newPass)
	if err != nil {
		return err.Error()
	}
	db, err := Connection()
	if err != nil {
		return fmt.Sprintf("error connecting to the database: %s", err)
	}

	hashedPassword, err := HashPass(newPass)
	if err != nil {
		return fmt.Sprintf("error hashing password: %s", err)
	}
	_, err = db.Exec("UPDATE users SET password = ? WHERE username = ?", hashedPassword, username)
	db.Close()
	if err != nil {
		return fmt.Sprintf("error updating password: %s", err)
	}
	return "Password updated successfully"
}