package api

import (
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/controller"
	"github.com/gorilla/mux"
	"github.com/aditya-411/mvc_assignment/pkg/middleware"
	"github.com/joho/godotenv"
)

func Start() {
	godotenv.Load("../../.env")
	router := mux.NewRouter()

	router.HandleFunc("/", controller.LandingPage_controller).Methods("GET")
	router.HandleFunc("/login", controller.Login).Methods("GET")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/register", controller.RegisterPage).Methods("GET")
	router.HandleFunc("/register", controller.RegisterPage).Methods("POST")

	router.HandleFunc("/user", controller.UserPage).Methods("GET")
	router.HandleFunc("/logout", controller.Logout).Methods("POST")

	router.HandleFunc("/user/browse", controller.BrowseBooks).Methods("GET")
	router.HandleFunc("/user/issue", controller.IssueBookPage).Methods("POST")
	router.HandleFunc("/user/issue_confirm", controller.ConfirmBookIssue).Methods("POST")
	router.HandleFunc("/user/my_books", controller.UserBooksPage).Methods("GET")

	router.Use(middleware.Middleware)


	http.ListenAndServe(":8000", router)
}
