package models

import(
	"github.com/aditya-411/mvc_assignment/pkg/types"
	//"time"
	//"math"
	//"fmt"
)


func AdminRequests() []types.User {
	db, err := Connection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT username FROM users WHERE admin_request = 1")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var requests []types.User
	for rows.Next() {
		request := types.User{}
		err := rows.Scan(&request.Username)
		if err != nil {
			panic(err)
		}
		requests = append(requests, request)
	}
	return requests
}

func ApproveAdminRequest(username string) string {
	db, err := Connection()
	if err != nil {
		return "error connecting to database"
	}
	defer db.Close()
	_, err = db.Exec("UPDATE users SET admin_request = 0, isadmin = 1 WHERE username = ?", username)
	if err != nil {
		return "error updating database"
	}
	return "Request Approved for " + username
}

func RejectAdminRequest(username string) string {
	db, err := Connection()
	if err != nil {
		return "error connecting to database"
	}
	defer db.Close()
	_, err = db.Exec("UPDATE users SET admin_request = 0 WHERE username = ?", username)
	if err != nil {
		return "error updating database"
	}
	return "Request Rejected for " + username
}