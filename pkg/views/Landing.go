package views

import (
	"html/template"
)

func LandingPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/index.html"))
	return temp
}
