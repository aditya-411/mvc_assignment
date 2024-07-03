package views

import (
	"net/http"
	"html/template"
	"github.com/aditya-411/mvc_assignment/pkg/types"
)

func UpdatePassView(w http.ResponseWriter, r *http.Request, message types.Message) {
	tmpl := template.Must(template.ParseFiles("templates/update_password.html"))
	tmpl.Execute(w, message)
}

