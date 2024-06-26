package api

import (
	"net/http"

	"github.com/aditya-411/mvc_assignment/pkg/controller"
	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Welcome).Methods("GET")
	r.HandleFunc("/add", controller.Add).Methods("POST")
	r.HandleFunc("/list", controller.List).Methods("GET")

	http.ListenAndServe(":8000", r)
}
