package main

import (
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/services"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/repos"
	BookRepo "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/repos/book"
	Router "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/router"
)

func main() {
	db := repos.New()
	// ------- Make Repository
	bookRepo := BookRepo.New(db)

	bookSrv := services.New(bookRepo)

	myRouter := Router.New(bookSrv)
	myRouter.Start()
}
