package controller

import (
	"net/http"

	//"github.com/aditya-411/mvc_assignment/pkg/models"
	"github.com/aditya-411/mvc_assignment/pkg/views"
)



func UserPage(writer http.ResponseWriter, request *http.Request) {
	t := views.UserPage()
	t.Execute(writer, nil)
}