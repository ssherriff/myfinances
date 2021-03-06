package model

import fieldbook "github.com/trexart/go-fieldbook"

const (
	transactionSheet = "transactions"
	categorySheet    = "categories"
)

//TransactionFieldbookDatastore Implementation of the datastore for Fieldbook
type TransactionFieldbookDatastore struct {
	Client *fieldbook.Client
}

//GetAllTransactions gets a list of all transactions stored in Fieldbook in the 'Transactions' sheet
func (ds *TransactionFieldbookDatastore) GetAllTransactions() ([]Transaction, error) {
	var transactions []Transaction
	var options fieldbook.QueryOptions

	err := ds.Client.ListRecords(transactionSheet, &transactions, &options)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

//GetAllTransactionsByMonthYear gets transactions for the month and year provided
func (ds *TransactionFieldbookDatastore) GetAllTransactionsByMonthYear(month, year int) ([]Transaction, error) {
	//Fieldbook doesn't have searching, for the sake of demo, just return all transactions
	// db or other datastore would be more capable
	return ds.GetAllTransactions()
}

//SaveTransaction create or update transaction record
func (ds *TransactionFieldbookDatastore) SaveTransaction(transaction *Transaction) (err error) {
	if transaction.ID == 0 {
		err = ds.Client.CreateRecord(transactionSheet, transaction)
	} else {
		err = ds.Client.UpdateRecord(transactionSheet, transaction.ID, transaction)
	}
	return
}

//GetAllCategories gets a list of all categories for grouping transactions stored in Fieldbook in the 'Transactions' sheet
func (ds *TransactionFieldbookDatastore) GetAllCategories() ([]Category, error) {
	var categories []Category
	var options fieldbook.QueryOptions

	err := ds.Client.ListRecords(categorySheet, &categories, &options)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
