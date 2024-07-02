package views

import (
	"html/template"
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/types"
	
)

func BookBrowsePage(w http.ResponseWriter, books types.ListBooks) {
	temp := template.Must(template.ParseFiles("templates/browse_books_user.html"))
	temp.Execute(w, books)
}

func IssueBookPage(w http.ResponseWriter, details types.IssueBookPage) {
	temp := template.Must(template.ParseFiles("templates/issue_book.html"))
	temp.Execute(w, details)
}