package controller

import (
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/types"
	"github.com/aditya-411/mvc_assignment/pkg/views"
	"github.com/aditya-411/mvc_assignment/pkg/models"
	"strconv"
)

func BookRequestsPage(w http.ResponseWriter, r *http.Request) {
	requests := models.BookRequests()
	details := types.BookRequestManagementPage{
		Requests: requests,
		Message:  "",
	}
	views.ManageRequestsPage(w, details)
}

func ApproveRequest(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err!=nil{
		panic(err)
	}
	message := models.ApproveBookRequest(id)
	requests := models.BookRequests()
	details := types.BookRequestManagementPage{
		Requests: requests,
		Message:  message,
	}
	views.ManageRequestsPage(w, details)
}

func RejectRequest(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err!=nil{
		panic(err)
	}
	message := models.DenyBookRequest(id)
	requests := models.BookRequests()
	details := types.BookRequestManagementPage{
		Requests: requests,
		Message:  message,
	}
	views.ManageRequestsPage(w, details)
}
