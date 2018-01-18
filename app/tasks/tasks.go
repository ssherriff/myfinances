package tasks

import "github.com/ssherriff/my-finances-app/app/model"

type Scheduler struct {
	TransactionDatastore model.TransactionDatastore
	ReoccuringDatastore  model.ReoccuringDatastore
}
