package controller

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/shopspring/decimal"
	fieldbook "github.com/trexart/go-fieldbook"

	"github.com/ssherriff/myfinances/app/model"
	"github.com/ssherriff/myfinances/app/view"
)

//Transaction Controller for transaction requests
type Transaction struct {
	TransactionDatastore model.TransactionDatastore
}

// Routes creates a REST router for the products resource
func (c *Transaction) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Get("/", c.List)    // GET /transactions - read a list of transactions
	r.Post("/", c.Create) // POST /transactions - create a transaction

	return r
}

//List Transaction list page
func (c *Transaction) List(w http.ResponseWriter, r *http.Request) {
	transactions, err := c.TransactionDatastore.GetAllTransactions()
	if err != nil {
		log.Println("Error getting transactions: ", err)
		return
	}

	categories, err := c.TransactionDatastore.GetAllCategories()
	if err != nil {
		log.Println("Error getting categories: ", err)
		return
	}

	var data = &struct {
		Transactions []model.Transaction
		Categories   []model.Category
	}{
		Transactions: transactions,
		Categories:   categories,
	}
	view.RenderTemplate(w, "transactions.html", view.NewPage("Transactions", "transactions", data))
}

//Create Transaction create
func (c *Transaction) Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	dateStr := r.FormValue("date")
	description := r.FormValue("description")
	categoryStr := r.FormValue("category")
	amountStr := r.FormValue("amount")

	// TODO Validate and error check

	date, _ := time.Parse("02/01/2006", dateStr)
	categoryID, _ := strconv.Atoi(categoryStr)
	amount, _ := decimal.NewFromString(amountStr)

	transaction := &model.Transaction{
		Date: fieldbook.Time{
			Time: date,
		},
		Description: description,
		Category: []model.Category{
			model.Category{ID: categoryID},
		},
		Amount: amount,
	}

	err := c.TransactionDatastore.SaveTransaction(transaction)

	if err != nil {
		errResponse := &ErrResponse{
			HTTPStatusCode: http.StatusInternalServerError,
			Err:            err,
			ErrorText:      "Error saving transaction",
		}
		errResponse.Render(w, r)
		return
	}
	render.Status(r, http.StatusOK)
}
