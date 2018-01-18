package model

import (
	"time"

	fieldbook "github.com/trexart/go-fieldbook"
)

const (
	reoccuringSheet = "reoccuring"
)

//ReoccuringFieldbookDatastore Implementation of the datastore for Fieldbook
type ReoccuringFieldbookDatastore struct {
	Client *fieldbook.Client
}

//GetAllReoccuring gets a list of all transactions stored in Fieldbook in the 'Transactions' sheet
func (ds *ReoccuringFieldbookDatastore) GetAllReoccuring() ([]Reoccuring, error) {
	var reoccuring []Reoccuring
	var options fieldbook.QueryOptions

	err := ds.Client.ListRecords(reoccuringSheet, &reoccuring, &options)
	if err != nil {
		return nil, err
	}
	return reoccuring, nil
}

//SaveReoccuring create or update a reoccuring record
func (ds *ReoccuringFieldbookDatastore) SaveReoccuring(reoccuring *Reoccuring) (err error) {
	if reoccuring.ID == 0 {
		err = ds.Client.CreateRecord(reoccuringSheet, reoccuring)
	} else {
		err = ds.Client.UpdateRecord(reoccuringSheet, reoccuring.ID, reoccuring)
	}
	return
}

//GetAllReoccuringOnOrBeforeDate gets a list of all transactions stored in Fieldbook in the 'Transactions' sheet
func (ds *ReoccuringFieldbookDatastore) GetAllReoccuringOnOrBeforeDate(date time.Time) ([]Reoccuring, error) {
	var reoccuring []Reoccuring
	var all []Reoccuring

	all, err := ds.GetAllReoccuring()
	if err != nil {
		return nil, err
	}

	for _, r := range all {
		if r.NextDate.Before(date) || r.NextDate.Equal(date) {
			reoccuring = append(reoccuring, r)
		}
	}

	return reoccuring, nil
}
