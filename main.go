package main

import (
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/services"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/repos"
	BookRepo "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/repos/book"
	MyRouter "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/routers/myrouter"
)

func main() {

	db := repos.New()
	// ------- Make Repository
	bookRepo := BookRepo.New(db)

	bookSrv := services.New(bookRepo)

	myRouter := MyRouter.New(bookSrv)

	myRouter.Start()

}
