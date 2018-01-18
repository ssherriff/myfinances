package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ssherriff/my-finances-app/app/model"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTransaction_List(t *testing.T) {
	Convey("Given a http request to /transactions ", t, func() {
		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest("GET", "/transactions", nil)
		if err != nil {
			t.Fatal(err)
		}
		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		ds := new(model.TestTransactionDatastore)
		ds.On("GetAllTransactions").Return([]model.Transaction{}, nil)
		transactionController := &Transaction{
			TransactionDatastore: ds,
		}
		handler := http.HandlerFunc(transactionController.List)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		Convey("Mock expectations should be met", func() {
			So(ds.AssertExpectations(t), ShouldBeTrue)
		})

		Convey("Http status should be 200", func() {
			So(rr.Code, ShouldEqual, http.StatusOK)
		})
	})
}
