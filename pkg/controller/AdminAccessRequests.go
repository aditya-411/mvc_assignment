package controller

import (
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/types"
	"github.com/aditya-411/mvc_assignment/pkg/views"
	"github.com/aditya-411/mvc_assignment/pkg/models"
	//"strconv"
)

func AdminAccessRequestsPage(w http.ResponseWriter, r *http.Request) {
	requests := models.AdminRequests()
	details := types.AdminAccessRequestsPage{
		Requests: requests,
		Message:  "",
	}
	views.AdminAccessRequestsPage(w, details)
}

func ApproveAdminRequest(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	message := models.ApproveAdminRequest(username)
	requests := models.AdminRequests()
	details := types.AdminAccessRequestsPage{
		Requests: requests,
		Message:  message,
	}
	views.AdminAccessRequestsPage(w, details)
}

func RejectAdminRequest(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	message := models.RejectAdminRequest(username)
	requests := models.AdminRequests()
	details := types.AdminAccessRequestsPage{
		Requests: requests,
		Message:  message,
	}
	views.AdminAccessRequestsPage(w, details)
}