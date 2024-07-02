package controller

import (
	"net/http"
)

func Logout(writer http.ResponseWriter, request *http.Request) {
	cookie := &http.Cookie{
		Name:   "login_token",
		Value:  "",
	}
	http.SetCookie(writer, cookie)
	http.Redirect(writer, request, "/", http.StatusSeeOther)
}