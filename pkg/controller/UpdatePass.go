package controller

import (
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/views"
	"github.com/aditya-411/mvc_assignment/pkg/models"
	"github.com/aditya-411/mvc_assignment/pkg/types"
)

func UpdatePassPage(w http.ResponseWriter, r *http.Request) {
	var message types.Message
	views.UpdatePassView(w, r, message)
}

func UpdatePass(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	currentPass := r.FormValue("currentPassword")
	newPass := r.FormValue("newPassword")
	confirmPass := r.FormValue("confirmPassword")
	user := r.Context().Value(types.Key("username")).(string)
	var message types.Message = types.Message{
		Text: models.UpdatePass(currentPass, newPass, confirmPass, user),
	}
	println(message.Text)
	views.UpdatePassView(w, r, message)
}