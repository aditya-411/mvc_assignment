package views

import (
	"html/template"
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/types"
	
)

func UpdateBookView(w http.ResponseWriter, details types.IssueBookPage) {
	temp := template.Must(template.ParseFiles("templates/update_book.html"))
	temp.Execute(w, details)
}