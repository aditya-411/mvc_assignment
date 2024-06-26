package controller

import (
	"net/http"

	"../models"
	"../views"
)

func List(writer http.ResponseWriter, request *http.Request) {
	// writer.WriteHeader(http.StatusOK)
	t := views.ListPage()
	writer.WriteHeader(http.StatusOK)
	booksList := models.FetchBooks()
	t.Execute(writer, booksList)
}
