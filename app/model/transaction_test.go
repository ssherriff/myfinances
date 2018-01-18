package model

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	fieldbook "github.com/trexart/go-fieldbook"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTransaction(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given a Transaction ", t, func() {
		amount, _ := decimal.NewFromString("-145.00")
		transaction := &Transaction{
			Date: fieldbook.Time{
				Time: time.Date(2018, 1, 20, 0, 0, 0, 0, time.UTC),
			},
			Amount: amount,
			Category: []Category{Category{
				Name: "Household",
			}},
		}

		Convey("When Date was set to 2018-01-20", func() {
			Convey("DateDisplay() should equal 20/01/2018", func() {
				So(transaction.DateDisplay(), ShouldEqual, "20/01/2018")
			})
		})

		Convey("When Amount was set to -145.000", func() {
			Convey("AmountDisplay() should equal -$145.00", func() {
				So(transaction.AmountDisplay(), ShouldEqual, "-$145.00")
			})
		})

		Convey("When category has 1 entry", func() {
			Convey("CategoryDisplay() should equal Household", func() {
				So(transaction.CategoryDisplay(), ShouldEqual, "Household")
			})
		})

		Convey("When category has 0 entries", func() {
			transaction.Category = []Category{}
			Convey("CategoryDisplay() should be empty", func() {
				So(transaction.CategoryDisplay(), ShouldBeZeroValue)
			})
		})

		Convey("When category has 2 entries", func() {
			transaction.Category = []Category{
				Category{
					Name: "Household",
				},
				Category{
					Name: "Personal",
				},
			}
			Convey("CategoryDisplay() should be empty", func() {
				So(transaction.CategoryDisplay(), ShouldBeZeroValue)
			})
		})
	})
}
