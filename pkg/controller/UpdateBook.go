package controller

import (
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/views"
	"github.com/aditya-411/mvc_assignment/pkg/types"
	"github.com/aditya-411/mvc_assignment/pkg/models"
)

func UpdateBookPage(w http.ResponseWriter, r *http.Request) {
	book := r.FormValue("title")
	author := r.FormValue("author")
	publisher := r.FormValue("publisher")
	quantity := r.FormValue("quantity")
	details := types.IssueBookPage{
		Book: types.Book{
			Name: book,
			Author: author,
			Publisher: publisher,
			Quantity: quantity,
		},
		Message: "",
	}
	views.UpdateBookView(w, details)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	old_book_name := r.FormValue("old_title")
	old_book_author := r.FormValue("old_author")
	old_book_publisher := r.FormValue("old_publisher")
	old_book_quantity := r.FormValue("old_quantity")
	bookName := r.FormValue("title")
	author := r.FormValue("author")
	publisher := r.FormValue("publisher")
	quantity := r.FormValue("quantity")
	book := types.Book{
		Name: bookName,
		Author: author,
		Publisher: publisher,
		Quantity: quantity,
	}
	message, updated := models.UpdateBook(old_book_name, book)
	if !updated{
		book = types.Book{
			Name: old_book_name,
			Author: old_book_author,
			Publisher: old_book_publisher,
			Quantity: old_book_quantity,
		}
	}
	details := types.IssueBookPage{
		Book: book,
		Message: message,
	}
	views.UpdateBookView(w, details)
}