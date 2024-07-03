package models

import (

	"github.com/aditya-411/mvc_assignment/pkg/types"
)

// functions to use in main functions


func BookIssued(book types.Book, user string) bool {
	db, err := Connection()
	if err != nil {
		panic(err)
	}

	var count int
	db.QueryRow("SELECT COUNT(username) FROM transactions WHERE username = ? AND title = ? and issued_at is NOT NULL and returned_at is NULL", user, book.Name).Scan(&count)	
	db.Close()
	return count > 0
}

func IssueRequestPending(user string) bool {
	db, err := Connection()
	if err != nil {
		panic(err)
	}
	var count int
	db.QueryRow("SELECT COUNT(username) FROM transactions WHERE username = ? AND issued_at is NULL AND returned_at is NULL", user).Scan(&count)
	db.Close()
	return count > 0
}


func MakeDetailsStructIssuePage(book types.Book, message string, show_button bool) types.IssueBookPage {
	return types.IssueBookPage{
		Book: book,
		Message: message,
		Show_button: show_button,
	}
}



// main function to return the details of the book to controller
func IssuePageDetails(book types.Book, user string) types.IssueBookPage {
	message := ""
	show_button := true
	if BookIssued(book, user) {
		message = "You already have this book issued"
		show_button = false
	} else if IssueRequestPending(user) {
		message = "You already have a pending issue request for some book"
		show_button = false
	}
	return MakeDetailsStructIssuePage(book, message, show_button)
}


// main function to confirm issue of the book
func IssueConfirm(book types.Book, user string) types.IssueBookPage {
	db, err := Connection()
	message := ""
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO transactions (username, title, request_status) VALUES (?, ?, '1')", user, book.Name)
	db.Close()
	if err != nil {
		 println(err.Error())
		 message = "Error in issuing book"
	} else {
		message = "Book issued request raised successfully"
	}

	return MakeDetailsStructIssuePage(book, message, false)
}