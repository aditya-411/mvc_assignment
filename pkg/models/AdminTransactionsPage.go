package models

import (
	"math"
	"strconv"
	"time"

	"github.com/aditya-411/mvc_assignment/pkg/types"
)


func GetAllCurrentlyIssuedBooks (allBooks types.ListBooks) []types.CurrentTransaction {
	db , err := Connection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	selectSql := "SELECT id, username, title, issued_at FROM transactions where returned_at is NULL and issued_at is not NULL"
	rows, err := db.Query(selectSql)

	if err != nil {
		panic(err)
	}

	var fetchBooks []types.CurrentTransaction
	for rows.Next() {
		var book types.CurrentTransaction
		var issued_at time.Time
		err := rows.Scan(&book.Id, &book.Username, &book.Name, &issued_at)
		if err != nil {
			panic(err)
		}
		book.IssueDate = issued_at.Format("02 Jan 2006")
		book.ReturnDate = issued_at.AddDate(0, 0, 7).Format("02 Jan 2006")
		book.Fine = strconv.FormatFloat(math.Max(0,math.Floor(time.Since(issued_at).Hours() / 24)-7) * 10, 'f', -1, 64)
		book.Author = getBookAuthor(book.Name, allBooks)
		fetchBooks = append(fetchBooks, book)
	}

	
	return fetchBooks

}



func AllPrevTransactions (books types.ListBooks) []types.PrevTransaction {
	db , err := Connection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	selectSql := "SELECT id, username, title, issued_at, returned_at FROM transactions where returned_at is NOT NULL"
	rows, err := db.Query(selectSql)

	if err != nil {
		panic(err)
	}

	var fetchBooks []types.PrevTransaction
	for rows.Next() {
		var book types.PrevTransaction
		var issued_at time.Time
		var returned_at time.Time
		err := rows.Scan(&book.Id, &book.Username, &book.Name, &issued_at, &returned_at)
		if err != nil {
			panic(err)
		}
		book.IssueDate = issued_at.Format("02 Jan 2006")
		book.ReturnDate = returned_at.Format("02 Jan 2006")
		book.Author = getBookAuthor(book.Name, books)

		fetchBooks = append(fetchBooks, book)
	}

	
	return fetchBooks

}


func CreateAdminTransactionsStruct(a []types.CurrentTransaction, c []types.PrevTransaction, message string) types.UserBooksPage {
	return types.UserBooksPage{
		CurrentlyIssuedBooks: a,
		PrevTransactions: c,
		Message: message,
	}
}
