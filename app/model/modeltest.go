package model

import "github.com/stretchr/testify/mock"

// Create a mock for other packages to use to test
/* USAGE:
ds := new(model.TestTransactionDatastore)
ds.On("GetAllTransactions").Return([]model.Transaction{}, nil)
*/
type TestTransactionDatastore struct {
	mock.Mock
}

func (ds *TestTransactionDatastore) GetAllTransactions() ([]Transaction, error) {
	args := ds.Called()
	return args.Get(0).([]Transaction), args.Error(1)
}

func (ds *TestTransactionDatastore) GetAllCategories() ([]Category, error) {
	args := ds.Called()
	return args.Get(0).([]Category), args.Error(1)
}
