package controller

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/ssherriff/my-finances-app/app/model"
	"github.com/ssherriff/my-finances-app/app/view"
)

//Transaction Controller for transaction requests
type Transaction struct {
	TransactionDatastore model.TransactionDatastore
}

// Routes creates a REST router for the products resource
func (c *Transaction) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Get("/", c.List) // GET /transactions - read a list of transactions

	return r
}

//List Transaction list page
func (c *Transaction) List(w http.ResponseWriter, r *http.Request) {
	transactions, err := c.TransactionDatastore.GetAllTransactions()
	if err != nil {
		log.Println("Error getting transactions: ", err)
		return
	}

	var data = &struct {
		Transactions []model.Transaction
	}{
		Transactions: transactions,
	}
	view.RenderTemplate(w, "transactions.html", view.NewPage("Transactions", "transactions", data))
}
