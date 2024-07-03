package views

import (
	"html/template"
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/types"
	
)

func AdminAccessRequestsPage(w http.ResponseWriter, details types.AdminAccessRequestsPage) {
	temp := template.Must(template.ParseFiles("templates/admin_access_requests.html"))
	temp.Execute(w, details)
}