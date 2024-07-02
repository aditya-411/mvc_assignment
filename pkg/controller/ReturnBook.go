package controller

import (
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/types"
	"github.com/aditya-411/mvc_assignment/pkg/models"
	//"github.com/aditya-411/mvc_assignment/pkg/types"
	"github.com/aditya-411/mvc_assignment/pkg/views"
)

func ReturnBook(w http.ResponseWriter, r *http.Request) {
	book := models.GetReturnBook(r)
	user := r.Context().Value(types.Key("username")).(string)
	message := models.ReturnBook(user, book)
	allBooks := models.FetchBooks()
	CurrentlyIssuedBooks := models.GetCurrentlyIssuedBooks("a", allBooks)
	PendingApprovals := models.PendingApprovals("a", allBooks)
	PrevTransactions := models.PrevTransactions("a", allBooks)
	details := models.CreateUserBooksPageStruct(CurrentlyIssuedBooks, PendingApprovals, PrevTransactions, message)
	views.UserBooksPage(w, details)
}