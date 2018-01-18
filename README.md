# My Finances

Simple personal budgeting application written in Go.

Currently in early stages and allows for keeping track of transactions and scheduling reoccuring transactions.

Only datastore available at the moment is Fieldbook. 

## Issues

  * Authentication
  * Transaction form validation
  * Transacton editing (inline?)
  * ReactJS UI
  * Replace http page loads in controllers with API for ReactJS
  * Replace Fieldbook with another datastore like Postgres
  * Creation of categories and more complex structures. Multiple levels.
  * Add concept of accounts to group transactions by Bank Account, Credit Cards, Investment Accounts, etc.
  * Allow for setting budgets for different categories
  * More options for reoccuring expenses: Yearly, Quarterly
  * Auto-budget reoccuring expenses, allowing for yearly and quarterly expenses to be divided into monthly budget amounts
  * Auto-import transactions from Banks or Files
  * Reconcile
  * Set savings goals to budget towards
  * Bills and reminders
  * Warnings, notification when over budget
  * ...

## Developers

You'll need to get the following packages:

    go get github.com/trexart/go-fieldbook
    go get github.com/go-chi/chi
    go get github.com/go-chi/render
    go get github.com/robfig/cron
    go get github.com/CloudyKit/jet
    go get github.com/leekchan/accounting
    go get github.com/shopspring/decimal

    go get github.com/smartystreets/goconvey
    go get github.com/stretchr/testify/mock

Build:

    go build -o myfinances

Run requires some parameters to be passed in. 

    ./myfinances -key=FIELDBOOKKEY -secret=FIELDBOOKSECRET -bookID=FIELDBOOKBOOKID

Easiest to use helper run script. Copy run.default.sh and just rename as run.sh.
Edit this line with your details:

    ./myfinances -key=key-1 -secret=SECRET -bookID=BOOKID

Save and just run as:

    ./run.sh

## Testing

App is setup to use GoConvey, but standard go testing can also be used. A helper script is available.

    ./goconvey.sh