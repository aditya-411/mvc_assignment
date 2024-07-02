package controller

import (
	"fmt"
	//"strconv"
	"net/http"

	"github.com/aditya-411/mvc_assignment/pkg/models"
	"github.com/aditya-411/mvc_assignment/pkg/views"
	//"github.com/aditya-411/mvc_assignment/pkg/types"
)

func UserBooksPage(w http.ResponseWriter, r *http.Request) {
	allBooks := models.FetchBooks()
	CurrentlyIssuedBooks := models.GetCurrentlyIssuedBooks("a", allBooks)
	PendingApprovals := models.PendingApprovals("a", allBooks)
	PrevTransactions := models.PrevTransactions("a", allBooks)
	details := models.CreateUserBooksPageStruct(CurrentlyIssuedBooks, PendingApprovals, PrevTransactions, "")
	fmt.Println(details)
	views.UserBooksPage(w, details)
}