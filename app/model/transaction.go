package model

import (
	"github.com/leekchan/accounting"
	"github.com/shopspring/decimal"
	fieldbook "github.com/trexart/go-fieldbook"
)

//TransactionDatastore access methods for transaction data
type TransactionDatastore interface {
	GetAllTransactions() ([]Transaction, error)

	GetAllCategories() ([]Category, error)
}

//Category allows for grouping of transactions
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//Transaction model of an account transaction
type Transaction struct {
	Date        fieldbook.Time  `json:"date"`
	Description string          `json:"description"`
	Category    []Category      `json:"category"`
	Amount      decimal.Decimal `json:"amount"`
}

//DateDisplay returns date formatted for display
func (t *Transaction) DateDisplay() string {
	return t.Date.Format("02/01/2006")
}

//CategoryDisplay returns category name for display
func (t *Transaction) CategoryDisplay() (category string) {
	if len(t.Category) == 1 {
		category = t.Category[0].Name
	}
	return
}

//AmountDisplay returns date formatted for display
func (t *Transaction) AmountDisplay() string {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	amount, _ := t.Amount.Float64()
	return ac.FormatMoneyFloat64(amount)
}
