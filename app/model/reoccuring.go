package model

import (
	"time"

	"github.com/shopspring/decimal"
	fieldbook "github.com/trexart/go-fieldbook"
)

//Reoccuring model of a reoccuring transaction
type Reoccuring struct {
	ID          int             `json:"id,omitempty"`
	Type        string          `json:"type,omitempty"`
	StartDate   fieldbook.Time  `json:"start_date,omitempty"`
	NextDate    fieldbook.Time  `json:"next_date,omitempty"`
	Period      string          `json:"period,omitempty"`
	Description string          `json:"description,omitempty"`
	Category    []Category      `json:"category,omitempty"`
	Amount      decimal.Decimal `json:"amount,omitempty"`
}

//SetNextDate set the next date for reoocurance based on current NextDate value
func (r *Reoccuring) SetNextDate() {
	switch r.Period {
	case "Weekly":
		r.NextDate.Time = r.NextDate.Time.AddDate(0, 0, 7)
	case "Monthly":
		r.NextDate.Time = r.NextDate.Time.AddDate(0, 1, 0)
	}
}

//ReoccuringDatastore access methods for transaction data
type ReoccuringDatastore interface {
	GetAllReoccuring() ([]Reoccuring, error)
	GetAllReoccuringOnOrBeforeDate(date time.Time) ([]Reoccuring, error)
	SaveReoccuring(reoccuring *Reoccuring) error
}
