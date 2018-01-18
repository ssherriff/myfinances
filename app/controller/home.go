package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/shopspring/decimal"

	"github.com/ssherriff/my-finances-app/app/model"
	"github.com/ssherriff/my-finances-app/app/view"
)

//Home Controller for requests
type Home struct {
	TransactionDatastore model.TransactionDatastore
}

//Main home page
func (c *Home) Main(w http.ResponseWriter, r *http.Request) {

	// Most of the below would probably be better done in a SQL query
	// Fieldbook is restrictive
	transactions, err := c.TransactionDatastore.GetAllTransactionsByMonthYear(1, 2018)
	if err != nil {
		log.Println("Error getting transactions: ", err)
		// TODO add error page
		return
	}
	categories, err := c.TransactionDatastore.GetAllCategories()
	if err != nil {
		log.Println("Error getting categories: ", err)
		// TODO add error page
		return
	}

	spendingMap := map[string]decimal.Decimal{}

	totalIncomeAmount := decimal.Zero

	for _, c := range categories {
		if c.Name != "Income" {
			spendingMap[c.Name] = decimal.Zero
		}
	}

	for _, t := range transactions {
		categoryName := t.Category[0].Name
		if categoryName != "Income" {
			amount := spendingMap[categoryName].Add(t.Amount.Neg())
			spendingMap[categoryName] = amount
		} else {
			totalIncomeAmount = totalIncomeAmount.Add(t.Amount)
		}
	}

	spendingLabels := make([]string, len(spendingMap)+1)
	spendingData := make([]string, len(spendingMap)+1)
	total := decimal.Zero

	i := 0
	for key, value := range spendingMap {
		spendingLabels[i] = key

		total = total.Add(value)
		spendingData[i] = value.String()

		i++
	}

	unallocatedAmount := totalIncomeAmount.Sub(total)
	spendingLabels[i] = "Unallocated"
	spendingData[i] = unallocatedAmount.String()

	var data = &struct {
		SpendingLabels string
		SpendingData   string
	}{
		SpendingLabels: "\"" + strings.Join(spendingLabels, "\",\"") + "\"",
		SpendingData:   strings.Join(spendingData, ","),
	}
	view.RenderTemplate(w, "home.html", view.NewPage("Overview", "", data))
}
