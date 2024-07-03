package controller

import (
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/views"
	"github.com/aditya-411/mvc_assignment/pkg/types"
)

func AdminPage(w http.ResponseWriter, r *http.Request) {
	user := types.User{
		Username: r.Context().Value(types.Key("username")).(string),
		IsAdmin:  r.Context().Value(types.Key("isadmin")).(bool),
	}
	views.AdminHomePage(w, user)
}