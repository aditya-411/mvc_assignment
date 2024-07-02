package controller

import (
	"net/http"
	//"github.com/aditya-411/mvc_assignment/pkg/models"
	"github.com/aditya-411/mvc_assignment/pkg/views"
	"github.com/aditya-411/mvc_assignment/pkg/types"
)



func UserPage(writer http.ResponseWriter, request *http.Request) {
	t := views.UserPage()
	user := types.User{
		Username: request.Context().Value(types.Key("username")).(string),
		IsAdmin:  request.Context().Value(types.Key("isadmin")).(bool),
	}
	t.Execute(writer, user)
}