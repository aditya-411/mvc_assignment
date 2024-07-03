package views

import (
	"html/template"
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/types"
	
)

func ManageRequestsPage(w http.ResponseWriter, details types.BookRequestManagementPage) {
	temp := template.Must(template.ParseFiles("templates/manage_book_requests.html"))
	temp.Execute(w, details)
}