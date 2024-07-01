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

	router.Use(middleware.Middleware)


	http.ListenAndServe(":8000", router)
}
