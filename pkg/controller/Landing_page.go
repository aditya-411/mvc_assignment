package controller

import (
	"net/http"

	"github.com/aditya-411/mvc_assignment/pkg/views"
)

func LandingPage_controller(writer http.ResponseWriter, request *http.Request) {
	t := views.LandingPage()
	t.Execute(writer, nil)
}
