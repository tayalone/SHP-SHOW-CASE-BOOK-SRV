package app

import (
	"fmt"
	"net/http"

	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/ports"
	routers "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/routers"
	RouteInitor "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/routers/init"
	"github.com/tayalone/SHP-SHOW-CASE-ESS-PKG/mylog"
)

/*Route is Route Composite*/
type Route struct {
	routers.Route
}

var ar Route

func iSayPing(c routers.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "pong",
	})
}

func myCustomMdw(c routers.Context) {
	fmt.Println("I println from myCustomMdw")
	c.Next()
}

/*New return router For Application */
func New(b ports.BookSrv) routers.Route {
	mylog.LogInfo("Holay I use my lovely PKG")
	myRouter := RouteInitor.Init("GIN", routers.Config{Port: 3000})
	myRouter.GET("/ping", myCustomMdw, iSayPing)
	myRouter.GET("/fiber", func(c routers.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Show Value from /fiber",
		})
	})

	v1 := myRouter.Group("/v1")

	v1.GET("/ping", myCustomMdw, iSayPing)

	myRouter.GET("/book/:id", func(ctx routers.Context) {
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

	ar.Route = myRouter
	return &ar
}
