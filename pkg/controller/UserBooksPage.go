package controller

import (
	//"strconv"
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/models"
	"github.com/aditya-411/mvc_assignment/pkg/views"
	"github.com/aditya-411/mvc_assignment/pkg/types"
)

func UserBooksPage(w http.ResponseWriter, r *http.Request) {
	allBooks := models.FetchBooks()
	user := r.Context().Value(types.Key("username")).(string)
	CurrentlyIssuedBooks := models.GetCurrentlyIssuedBooks(user, allBooks)
	PendingApprovals := models.PendingApprovals(user, allBooks)
	PrevTransactions := models.PrevTransactions(user, allBooks)
	details := models.CreateUserBooksPageStruct(CurrentlyIssuedBooks, PendingApprovals, PrevTransactions, "")
	views.UserBooksPage(w, details)
}