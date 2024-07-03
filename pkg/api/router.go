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

	
	router.HandleFunc("/request_admin", controller.RequestAdminAccess).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("GET")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/register", controller.RegisterPage).Methods("GET")
	router.HandleFunc("/register", controller.RegisterPage).Methods("POST")
	router.HandleFunc("/update_password", controller.UpdatePassPage).Methods("GET")
	router.HandleFunc("/update_password", controller.UpdatePass).Methods("POST")
	router.HandleFunc("/logout", controller.Logout).Methods("POST")


	router.HandleFunc("/admin", controller.AdminPage).Methods("GET")
	router.HandleFunc("/admin/books", controller.BookCatalogue).Methods("GET")
	router.HandleFunc("/admin/add", controller.AddBook).Methods("POST")
	router.HandleFunc("/admin/remove", controller.RemoveBook).Methods("POST")
	router.HandleFunc("/admin/update", controller.UpdateBookPage).Methods("POST")
	router.HandleFunc("/admin/update_confirm", controller.UpdateBook).Methods("POST")
	router.HandleFunc("/admin/requests", controller.BookRequestsPage).Methods("GET")
	router.HandleFunc("/admin/requests/approve", controller.ApproveRequest).Methods("POST")
	router.HandleFunc("/admin/requests/deny", controller.RejectRequest).Methods("POST")
	router.HandleFunc("/admin/access", controller.AdminAccessRequestsPage).Methods("GET")
	router.HandleFunc("/admin/access/approve", controller.ApproveAdminRequest).Methods("POST")
	router.HandleFunc("/admin/access/deny", controller.RejectAdminRequest).Methods("POST")

	router.HandleFunc("/user", controller.UserPage).Methods("GET")
	router.HandleFunc("/user/browse", controller.BrowseBooks).Methods("GET")
	router.HandleFunc("/user/issue", controller.IssueBookPage).Methods("POST")
	router.HandleFunc("/user/issue_confirm", controller.ConfirmBookIssue).Methods("POST")
	router.HandleFunc("/user/my_books", controller.UserBooksPage).Methods("GET")
	router.HandleFunc("/user/return", controller.ReturnBook).Methods("POST")

	router.Use(middleware.Middleware)


	http.ListenAndServe(":8000", router)
}
