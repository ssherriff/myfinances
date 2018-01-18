package tasks

import "github.com/ssherriff/my-finances-app/app/model"

//Scheduler task scheduler
type Scheduler struct {
	TransactionDatastore model.TransactionDatastore
	ReoccuringDatastore  model.ReoccuringDatastore
}
