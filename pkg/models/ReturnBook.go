package models

import (
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/types"
)

func GetReturnBook(r *http.Request) types.Book {
	book := types.Book{
		Name: r.FormValue("title"),
	}
	return book
}

func BookReturnRequested(book types.Book, user string) bool {
	db, err := Connection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var count int
	db.QueryRow("SELECT COUNT(username) FROM transactions WHERE username = ? AND title = ? AND issued_at is NOT NULL AND returned_at is NULL AND request_status='-1'", user, book.Name).Scan(&count)
	return count > 0
}

func ReturnBook(user string, book types.Book) string {
	if !BookIssued(book, user){
		return "You don't have the book issued. Don't try to be smart."
	}
	if BookReturnRequested(book, user) {
		return "You have already requested for return of this book. Please wait for approval."
	}
	db, err := Connection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("UPDATE transactions SET request_status='-1' WHERE username = ? AND title = ? AND issued_at is NOT NULL AND returned_at is NULL", user, book.Name)
	if err != nil {
		return "Some error occurred. Please try again."
	}
	return "Book return request raised successfully. Please wait for approval."
}