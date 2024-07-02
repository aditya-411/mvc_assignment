package views

import (
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/types"
	"html/template"
)

func UserBooksPage(w http.ResponseWriter, userbooks types.UserBooksPage) {
	temp := template.Must(template.ParseFiles("templates/view_books_user.html"))
	temp.Execute(w, userbooks)
}