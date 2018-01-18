package tasks

import "github.com/ssherriff/myfinances/app/model"

//Scheduler task scheduler
type Scheduler struct {
	TransactionDatastore model.TransactionDatastore
	ReoccuringDatastore  model.ReoccuringDatastore
}
