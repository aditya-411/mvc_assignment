package models

import (
	"fmt"
)

func RemoveBook(bookName string) string {
	db, err := Connection()
	if err != nil {
		return fmt.Sprintf("error %s connecting to the database", err)
	}
	defer db.Close()
	var count int
	db.QueryRow("SELECT COUNT(title) FROM books WHERE title = ?", bookName).Scan(&count)	
	if count == 0{
		return fmt.Sprintf("Book %s does not exist in database", bookName)
	}

	deleteSql := "DELETE FROM transactions WHERE title = ?"
	_, err = db.Exec(deleteSql, bookName)
	if err != nil {
		return fmt.Sprintf("error %s error deleting transactions with this book from database", err)
	} 
	
	deleteSql = "DELETE FROM books WHERE title = ?"
	_, err = db.Exec(deleteSql, bookName)
	if err != nil {
		return fmt.Sprintf("error %s error deleting the book from database", err)
	} else {
		return fmt.Sprintf("successfully deleted %s from the database", bookName)
	}
}
