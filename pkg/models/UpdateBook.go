package models

import (
	"fmt"
	"strconv"

	"github.com/aditya-411/mvc_assignment/pkg/types"
)

func UpdateBook(bookName string, book types.Book) (string, bool) {
	db, err := Connection()
	if err != nil {
		return fmt.Sprintf("error %s connecting to the database", err), false
	}
	defer db.Close()
	var count int
	db.QueryRow("SELECT COUNT(title) FROM books WHERE title = ?", bookName).Scan(&count)	
	if count == 0{
		return fmt.Sprintf("Book %s does not exist in database", bookName), false
	}

	if book.Name == "" || book.Author == "" || book.Publisher == "" {
		return "One or more fields in the book details are empty", false
	}

	_ , err = strconv.Atoi(book.Quantity)
	if err != nil {
		return fmt.Sprintf("error %s converting quantity to integer", err), false
	}
	updateSql := "UPDATE books SET title = ? , author = ? , publisher = ?, quantity_left = ? WHERE title = ?"
	_, err = db.Exec(updateSql, book.Name, book.Author, book.Publisher, book.Quantity, bookName)
	if err != nil {
		return fmt.Sprintf("%s error updating the book from database", err), false
	} else {
		return fmt.Sprintf("successfully updated %s in the database", bookName), true
	}
}
