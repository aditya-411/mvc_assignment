package controller

import (
	// "fmt"
	"net/http"

	"github.com/aditya-411/mvc_assignment/pkg/models"
	"github.com/aditya-411/mvc_assignment/pkg/types"
	// "github.com/aditya-411/mvc_assignment/pkg/types"
	"github.com/aditya-411/mvc_assignment/pkg/views"
	// "golang.org/x/crypto/bcrypt"
)



func Login(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		// Process the login form data
		username := request.FormValue("username")
		password := request.FormValue("password")

		match , err:= models.ComparePass(username, password)
		if !match || err != nil {
			t := views.LoginPage()
			text := types.Message{Text: "Invalid username or password"}
			t.Execute(writer, text)
			return
		}
		
		// Create JWT token
		token, err := models.GenerateJWT(username)
		if err != nil {
			// Handle error
			http.Error(writer, "Failed to generate JWT token", http.StatusInternalServerError)
			return
		}

		// Set JWT token as a cookie
		cookie := &http.Cookie{
			Name:     "login_token",
			Value:    token,
			HttpOnly: true,
		}
		http.SetCookie(writer, cookie)
		http.Redirect(writer, request, "/user", http.StatusFound)
		return
		
		
		
		
	} else {
		// Render the login page for GET requests
		t := views.LoginPage()
		t.Execute(writer, nil)
	} 
}
