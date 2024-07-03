package views

import (
	"html/template"
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/types"
	
)

func AdminHomePage(w http.ResponseWriter, user types.User) {
	temp := template.Must(template.ParseFiles("templates/admin_home.html"))
	temp.Execute(w, user)
}