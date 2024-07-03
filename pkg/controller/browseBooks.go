package controller

import (
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/models"
	"github.com/aditya-411/mvc_assignment/pkg/views"
)

func BrowseBooks(w http.ResponseWriter, r *http.Request) {
	books := models.FetchBooks()
	views.BookBrowsePage(w, books)
}