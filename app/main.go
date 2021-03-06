package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/robfig/cron"
	fieldbook "github.com/trexart/go-fieldbook"

	"github.com/ssherriff/myfinances/app/controller"
	"github.com/ssherriff/myfinances/app/model"
	"github.com/ssherriff/myfinances/app/tasks"
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
	reoccuringDatastore := &model.ReoccuringFieldbookDatastore{
		Client: fieldbookClient,
	}

	// Controllers available to the application
	homeController := &controller.Home{
		TransactionDatastore: transactionDatastore,
	}
	transactionController := &controller.Transaction{
		TransactionDatastore: transactionDatastore,
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", homeController.Main)

	r.Mount("/transactions", transactionController.Routes())

	taskScheduler := &tasks.Scheduler{
		TransactionDatastore: transactionDatastore,
		ReoccuringDatastore:  reoccuringDatastore,
	}

	log.Println("Scheduling tasks")
	c := cron.New()
	c.AddFunc("@daily", func() {
		taskScheduler.RegisterReoccuring()
	})
	c.Start()

	log.Println("Running at localhost:3000")
	http.ListenAndServe(":3000", r)
}
