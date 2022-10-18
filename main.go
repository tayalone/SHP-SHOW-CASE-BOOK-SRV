package main

import (
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/services"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/repos"
	BookRepo "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/repos/book"
	router "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/routers"
)

func main() {
	db := repos.New()
	// ------- Make Repository
	bookRepo := BookRepo.New(db)

	bookSrv := services.New(bookRepo)

	myRouter := router.New(bookSrv)

	myRouter.Start()

}
