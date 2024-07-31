package models

import(
	"github.com/aditya-411/mvc_assignment/pkg/types"
	"time"
	"math"
	"fmt"
)


func BookRequests() []types.PendingApproval {
	db, err := Connection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, username, title, request_status, issued_at FROM transactions WHERE request_status!='0'")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var requests []types.PendingApproval
	for rows.Next() {
		var request types.PendingApproval
		var request_status string
		var issued_at *time.Time
		err := rows.Scan(&request.Id, &request.Username, &request.Name, &request_status, &issued_at)
		if err != nil {
			panic(err)
		}

		if request_status == "1" {
			request.Type = "Issue"
			request.Fine = 0
		} else {
			request.Type = "Return"
			request.Fine = int(math.Max(0,math.Floor(time.Since(*issued_at).Hours() / 24)-7) * 10)
		}

		requests = append(requests, request)
	}
	return requests
}


func ApproveBookRequest(id int) string {
	db, err := Connection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var Request string
	var book string
	var user string
	err = db.QueryRow("SELECT request_status, title, username FROM transactions WHERE id=?", id).Scan(&Request, &book, &user)
	if err != nil {
		return "There is no request with this id"
	}
	if Request == "0"{
		return "There is no request with this id"
	} else if Request == "1" {
		if BookNotLeft(types.Book{Name: book}){
			_, err = db.Exec("DELETE FROM transactions where id=?", id)
			if err != nil {
				return "error processing request"
			}
			return "Book not available"
		}
		
		var title string
		db.QueryRow("SELECT title FROM transactions WHERE id=?", id).Scan(&title)
		_, err = db.Exec("UPDATE books SET quantity_left=quantity_left - 1 where title = ?", title)
		if err != nil {
			return "error approving request 1"
		}
		_, err = db.Exec("UPDATE transactions SET request_status='0', issued_at=CURDATE() WHERE id=?", id)
		if err != nil {
			return "error approving request 2"
		}
		return fmt.Sprintf("%s issued successfully to %s", book, user)
	} else {
		var title string
		db.QueryRow("SELECT title FROM transactions WHERE id=?", id).Scan(&title)
		_, err = db.Exec("UPDATE books SET quantity_left=quantity_left + 1 where title = ?", title)
		if err != nil {
			return "error approving request 2"
		}
		_, err = db.Exec("UPDATE transactions SET request_status='0', returned_at=CURDATE() WHERE id=?", id)
		if err != nil {
			return "error approving request"
		}
		return fmt.Sprintf("%s returned successfully by %s", book, user)
	}
}


func DenyBookRequest(id int) string {
	db, err := Connection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var Request string
	var book string
	var user string
	err = db.QueryRow("SELECT request_status, title, username FROM transactions WHERE id=?", id).Scan(&Request, &book, &user)
	if err != nil {
		return "There is no request with this id"
	}
	if Request == "0"{
		return "There is no request with this id"
	} else if Request == "1" {
		_, err = db.Exec("DELETE FROM transactions WHERE id=?", id)
		if err != nil {
			return "error denying request"
		}
		return fmt.Sprintf("%s issue request to %s denied successfully", book, user)
	} else {
		_, err = db.Exec("UPDATE transactions SET request_status='0' WHERE id=?", id)
		if err != nil {
			return "error denying request"
		}
		return fmt.Sprintf("%s return request by %s denied successfully", book, user)
	}
}