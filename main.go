package main

import (
	"fmt"
	"net/http"

	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/ports"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/services"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/repos"
	BookRepo "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/repos/book"
	router "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/routers"
	RouteInitor "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/routers/init"
	"github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/mylog"
)

func iSayPing(c router.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "pong",
	})
}

func myCustomMdw(c router.Context) {
	fmt.Println("I println from myCustomMdw")
	c.Next()
}

type tmpRoute struct {
	router.Route
}

var mtr tmpRoute

func newRoute(b ports.BookSrv) router.Route {
	mylog.LogInfo("Holay I use my lovely PKG")
	myRouter := RouteInitor.Init("GIN", router.Config{Port: 3000})
	myRouter.GET("/ping", myCustomMdw, iSayPing)
	myRouter.GET("/fiber", func(c router.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Show Value from /fiber",
		})
	})

	v1 := myRouter.Group("/v1")

	v1.GET("/ping", myCustomMdw, iSayPing)

	myRouter.GET("/book/:id", func(ctx router.Context) {
		type getIDUri struct {
			ID uint `uri:"id" binding:"required"`
		}

		var gi getIDUri
		if err := ctx.BindURI(&gi); err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"msg": err.Error(),
			})
			return
		}
		b, errGetPk := b.GetByID(gi.ID)
		if errGetPk != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"msg": errGetPk.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "OK",
			"book":    b,
		})
	})

	mtr.Route = myRouter
	return &mtr
}

func main() {
	db := repos.New()
	// ------- Make Repository
	bookRepo := BookRepo.New(db)

	bookSrv := services.New(bookRepo)

	// // myRouter := MyRouter.New(bookSrv)

	// // myRouter.Start()
	// mySw := SwitchRouter.New(bookSrv)
	// mySw.Start()

	// myRouter := RouteInitor.Init("GIN", router.Config{Port: 3000})
	// myRouter.GET("/ping", myCustomMdw, iSayPing)
	// myRouter.GET("/fiber", func(c router.Context) {
	// 	c.JSON(http.StatusOK, map[string]interface{}{
	// 		"message": "Show Value from /fiber",
	// 	})
	// })

	// v1 := myRouter.Group("/v1")

	// v1.GET("/ping", myCustomMdw, iSayPing)

	// myRouter.Start()
	myRouter := newRoute(bookSrv)
	myRouter.Start()
}
