package views

import (
	"html/template"
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/types"
	
)

func AdminTransactionsPage(w http.ResponseWriter, details types.UserBooksPage) {
	temp := template.Must(template.ParseFiles("templates/admin_transactions.html"))
	temp.Execute(w, details)
}