package controller

import (
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/types"
	"github.com/aditya-411/mvc_assignment/pkg/models"
	//"github.com/aditya-411/mvc_assignment/pkg/types"
	"github.com/aditya-411/mvc_assignment/pkg/views"
)

func RequestAdminAccess(w http.ResponseWriter, r *http.Request) {
	message := models.RequestAdminAccess(r.Context().Value(types.Key("username")).(string))
	Message := types.User{
		Message: message,
		Username: r.Context().Value(types.Key("username")).(string),}
	page := views.UserPage()
	page.Execute(w, Message)
}