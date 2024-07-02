package controller

import (
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/models"
	"github.com/aditya-411/mvc_assignment/pkg/views"
	"github.com/aditya-411/mvc_assignment/pkg/types"
)

func BrowseBooks(w http.ResponseWriter, r *http.Request) {
	books := models.FetchBooks()
	views.BookBrowsePage(w, books)
}


func IssueBookPage(w http.ResponseWriter, r *http.Request) {
	book := models.GetDetailsBook(r)
	details := models.IssuePageDetails(book, r.Context().Value(types.Key("username")).(string))
	views.IssueBookPage(w, details)
}

func ConfirmBookIssue(w http.ResponseWriter, r *http.Request) {
	book := models.GetDetailsBook(r)
	user := r.Context().Value(types.Key("username")).(string)
	details := models.IssueConfirm(book, user)
	views.IssueBookPage(w, details)
}