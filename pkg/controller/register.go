package controller

import (
	"net/http"

	//"github.com/aditya-411/mvc_assignment/pkg/models"
	"github.com/aditya-411/mvc_assignment/pkg/views"
)



func RegisterPage(writer http.ResponseWriter, request *http.Request) {
	t := views.RegisterPage()
	t.Execute(writer, nil)
}
