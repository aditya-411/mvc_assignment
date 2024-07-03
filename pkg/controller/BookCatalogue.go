package controller

import (
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/views"
	"github.com/aditya-411/mvc_assignment/pkg/types"
	"github.com/aditya-411/mvc_assignment/pkg/models"
)

func BookCatalogue(w http.ResponseWriter, r *http.Request) {
	books := models.FetchBooks().Books
	details := types.BookCatalogue{
		Books: books,
		Message: "",
	}
	views.BookCatalogue(w, details)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	message := models.AddBook(r.FormValue("title"), r.FormValue("author"), r.FormValue("publisher"))
	books := models.FetchBooks().Books
	details := types.BookCatalogue{
		Books: books,
		Message: message,
	}
	views.BookCatalogue(w, details)
}