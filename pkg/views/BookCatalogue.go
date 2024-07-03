package views

import (
	"html/template"
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/types"
	
)

func BookCatalogue(w http.ResponseWriter, details types.BookCatalogue) {
	temp := template.Must(template.ParseFiles("templates/book_catalogue.html"))
	temp.Execute(w, details)
}

