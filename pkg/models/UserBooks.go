package models

import (
	"math"
	"strconv"
	"time"

	"github.com/aditya-411/mvc_assignment/pkg/types"
)

func getBookAuthor (bookName string, books types.ListBooks) string {
	for _, book := range books.Books {
		if book.Name == bookName {
			return book.Author
		}
	}
	return ""
}

func GetCurrentlyIssuedBooks (username string, allBooks types.ListBooks) []types.CurrentTransaction {
	db , err := Connection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	selectSql := "SELECT id, title, issued_at FROM transactions where username = ? and returned_at is NULL and issued_at is not NULL and request_status='0'"
	rows, err := db.Query(selectSql, username)

	if err != nil {
		panic(err)
	}

	var fetchBooks []types.CurrentTransaction
	for rows.Next() {
		var book types.CurrentTransaction
		var issued_at time.Time
		err := rows.Scan(&book.Id, &book.Name, &issued_at)
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


func PendingApprovals (username string, books types.ListBooks) []types.PendingApproval {
	db , err := Connection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	selectSql := "SELECT title, request_status FROM transactions where username = ? and request_status!='0'"
	rows, err := db.Query(selectSql, username)

	if err != nil {
		panic(err)
	}

	var fetchBooks []types.PendingApproval
	for rows.Next() {
		var book types.PendingApproval
		var request_status string
		err := rows.Scan(&book.Name, &request_status)
		if err != nil {
			panic(err)
		}
		if request_status == "1" {
			book.Type = "Issue"
		} else {
			book.Type = "Return"
		}

		book.Author = getBookAuthor(book.Name, books)

		fetchBooks = append(fetchBooks, book)
	}

	
	return fetchBooks

}


func PrevTransactions (username string, books types.ListBooks) []types.PrevTransaction {
	db , err := Connection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	selectSql := "SELECT id, title, issued_at, returned_at FROM transactions where username = ? AND returned_at is NOT NULL"
	rows, err := db.Query(selectSql, username)

	if err != nil {
		panic(err)
	}

	var fetchBooks []types.PrevTransaction
	for rows.Next() {
		var book types.PrevTransaction
		var issued_at time.Time
		var returned_at time.Time
		err := rows.Scan(&book.Id, &book.Name, &issued_at, &returned_at)
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


func CreateUserBooksPageStruct(a []types.CurrentTransaction, b []types.PendingApproval, c []types.PrevTransaction, message string) types.UserBooksPage {
	return types.UserBooksPage{
		CurrentlyIssuedBooks: a,
		PendingApprovals: b,
		PrevTransactions: c,
		Message: message,
	}
}
