package controller

import (
	"net/http"
	"github.com/aditya-411/mvc_assignment/pkg/views"
	"github.com/aditya-411/mvc_assignment/pkg/models"
)

func AdminTransactionsPage(w http.ResponseWriter, r *http.Request) {
	allBooks := models.FetchBooks()
	CurrentlyIssuedBooks := models.GetAllCurrentlyIssuedBooks(allBooks)
	PrevTransactions := models.AllPrevTransactions(allBooks)
	details := models.CreateAdminTransactionsStruct(CurrentlyIssuedBooks, PrevTransactions, "")
	views.AdminTransactionsPage(w, details)
}