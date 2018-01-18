package tasks

import (
	"log"
	"time"

	"github.com/ssherriff/my-finances-app/app/model"
)

func (task *Scheduler) RegisterReoccuring() {
	log.Println("Registering reoccuring tasks")
	reoccuring, err := task.ReoccuringDatastore.GetAllReoccuringOnOrBeforeDate(time.Now())
	if err != nil {
		log.Println("Error getting reoccuring transactions: ", err)
		return
	}

	for _, r := range reoccuring {
		// create new transaction
		transaction := &model.Transaction{
			Date:        r.NextDate,
			Description: r.Description,
			Category:    r.Category,
			Amount:      r.Amount,
		}
		err = task.TransactionDatastore.SaveTransaction(transaction)
		if err != nil {
			log.Println("Error saving new transaction: ", err)
		} else {
			r.SetNextDate()
			newR := &model.Reoccuring{
				ID:       r.ID,
				NextDate: r.NextDate,
			}
			err = task.ReoccuringDatastore.SaveReoccuring(newR)
			if err != nil {
				log.Println("Error updating reoccuring: ", err)
			}
		}
	}
}
