package controller

import (
	"net/http"

	"github.com/aditya-411/mvc_assignment/pkg/models"
	"github.com/aditya-411/mvc_assignment/pkg/views"
)

func List(writer http.ResponseWriter, request *http.Request) {
	// writer.WriteHeader(http.StatusOK)
	t := views.ListPage()
	writer.WriteHeader(http.StatusOK)
	booksList := models.FetchBooks()
	t.Execute(writer, booksList)
}
