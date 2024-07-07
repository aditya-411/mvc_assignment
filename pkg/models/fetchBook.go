package models

import (
	"fmt"
	"github.com/aditya-411/mvc_assignment/pkg/types"
	"net/http"
)

func FetchBooks() types.ListBooks {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}
	defer db.Close()
	selectSql := "SELECT * FROM books"
	rows, err := db.Query(selectSql)

	if err != nil {
		fmt.Printf("error %s querying the database", err)
	}

	var fetchBooks []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.Name, &book.Author, &book.Publisher)
		if err != nil {
			fmt.Printf("error %s scanning the row", err)
		}
		fetchBooks = append(fetchBooks, book)
	}

	var listBooks types.ListBooks
	listBooks.Books = fetchBooks
	return listBooks

}

func GetDetailsBook(request *http.Request) types.Book {
	book := types.Book{
		Name: request.FormValue("title"),
		Publisher: request.FormValue("publisher"),
		Author: request.FormValue("author"),
	}
	return book
}
