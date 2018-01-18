#!/bin/sh

go get github.com/trexart/go-fieldbook
go get github.com/go-chi/chi
go get github.com/go-chi/render
go get github.com/robfig/cron
go get github.com/CloudyKit/jet
go get github.com/leekchan/accounting
go get github.com/shopspring/decimal

go get github.com/smartystreets/goconvey
go get github.com/stretchr/testify/mock

cd app
go build -o myfinances
./myfinances -key=key-1 -secret=SECRET -bookID=BOOKID