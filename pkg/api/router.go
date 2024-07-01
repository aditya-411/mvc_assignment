package api

import (
	"net/http"

	"github.com/aditya-411/mvc_assignment/pkg/controller"
	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.LandingPage_controller).Methods("GET")
	r.HandleFunc("/login", controller.Login).Methods("GET")
	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/register", controller.RegisterPage).Methods("GET")
	r.HandleFunc("/user", controller.UserPage).Methods("GET")

	http.ListenAndServe(":8000", r)
}
