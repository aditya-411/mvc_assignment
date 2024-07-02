package controller

import (
	//"strconv"
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/models"
	"github.com/aditya-411/mvc_assignment/pkg/views"
)

func UserBooksPage(w http.ResponseWriter, r *http.Request) {
	allBooks := models.FetchBooks()
	CurrentlyIssuedBooks := models.GetCurrentlyIssuedBooks("a", allBooks)
	PendingApprovals := models.PendingApprovals("a", allBooks)
	PrevTransactions := models.PrevTransactions("a", allBooks)
	details := models.CreateUserBooksPageStruct(CurrentlyIssuedBooks, PendingApprovals, PrevTransactions, "")
	views.UserBooksPage(w, details)
}