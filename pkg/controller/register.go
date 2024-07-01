package controller

import (
	"net/http"

	"github.com/aditya-411/mvc_assignment/pkg/models"
	"github.com/aditya-411/mvc_assignment/pkg/views"
	"github.com/aditya-411/mvc_assignment/pkg/types"
)



func RegisterPage(writer http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		t := views.RegisterPage()
		t.Execute(writer, nil)
		return 
	}

	username := request.FormValue("username")
	password := request.FormValue("password")
	confirmPassword := request.FormValue("confirm_password")

	err := models.Registration(username, password, confirmPassword)
	t := views.RegisterPage()
	if err != nil {
		message := types.Message{Text: err.Error()}
		t.Execute(writer, message)
		return
	}
	message := types.Message{Text: "Registration successful"}
	t.Execute(writer, message)


}
