package main

import (
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/services"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/repos"
	BookRepo "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/repos/book"
	App "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/routers/app"
)

func main() {
	db := repos.New()
	// ------- Make Repository
	bookRepo := BookRepo.New(db)

	bookSrv := services.New(bookRepo)

	myRouter := App.New(bookSrv)
	myRouter.Start()
}
