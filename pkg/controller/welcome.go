package controller

import (
	"net/http"

	"github.com/aditya-411/mvc_assignment/pkg/views"
)

func Welcome(writer http.ResponseWriter, request *http.Request) {
	t := views.StartPage()
	t.Execute(writer, nil)
}
