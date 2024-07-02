package models

import (
	"net/http"

	"github.com/aditya-411/mvc_assignment/pkg/types"
)

func GetDetailsBook(request *http.Request) types.Book {
	book := types.Book{
		Name: request.FormValue("title"),
		Publisher: request.FormValue("publisher"),
		Author: request.FormValue("author"),
	}
	return book
}


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


func MakeDetailsStruct(book types.Book, message string, show_button bool) types.IssueBookPage {
	return types.IssueBookPage{
		Book: book,
		Message: message,
		Show_button: show_button,
	}
}

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
	return MakeDetailsStruct(book, message, show_button)
}