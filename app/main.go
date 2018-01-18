package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/robfig/cron"
	fieldbook "github.com/trexart/go-fieldbook"

	"github.com/ssherriff/my-finances-app/app/controller"
	"github.com/ssherriff/my-finances-app/app/model"
	"github.com/ssherriff/my-finances-app/app/tasks"
	"github.com/ssherriff/my-finances-app/app/view"
)

func main() {

	keyPtr := flag.String("key", "", "Key for fieldbook API")
	secretPtr := flag.String("secret", "", "Secret for fieldbook API")
	bookIDPtr := flag.String("bookID", "", "Book ID for fieldbook API")
	flag.Parse()

	fieldbookClient := fieldbook.NewClient(*keyPtr, *secretPtr, *bookIDPtr)

	// Datasources available to the application
	transactionDatastore := &model.TransactionFieldbookDatastore{
		Client: fieldbookClient,
	}

	// Controllers available to the application
	transactionController := &controller.Transaction{
		TransactionDatastore: transactionDatastore,
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		view.RenderTemplate(w, "home.html", view.NewPage("Home", "", nil))
	})

	r.Mount("/transactions", transactionController.Routes())

	log.Println("Scheduling RegisterReoccuring task")
	c := cron.New()
	c.AddFunc("@daily", func() {
		tasks.RegisterReoccuring()
	})
	c.Start()

	log.Println("Running at localhost:3000")
	http.ListenAndServe(":3000", r)
}
