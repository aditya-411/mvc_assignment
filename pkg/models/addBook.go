package models

import (
	"fmt"
)

func AddBook(bookName string, author string, publisher string) string {
	db, err := Connection()
	if err != nil {
		return fmt.Sprintf("error %s connecting to the database", err)
	}
	defer db.Close()
	var count int
	db.QueryRow("SELECT COUNT(title) FROM books WHERE title = ?", bookName).Scan(&count)	
	if count > 1{
		return fmt.Sprintf("Book %s already exists in the database", bookName)
	}

	if bookName == "" || author == "" || publisher == "" {
		return "One or more fields in the book details are empty"
	}

	insertSql := "INSERT INTO books (title, author, publisher) VALUES (?, ?, ?)"
	_, err = db.Exec(insertSql, bookName, author, publisher)
	if err != nil {
		return fmt.Sprintf("error %s inserting into the database", err)
	} else {
		return fmt.Sprintf("successfully inserted %s into the database", bookName)
	}
}

