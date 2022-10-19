package init

import (
	routers "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/routers"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/routers/fiber"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/routers/gin"
)

/*Init Reouter Router Instant */
func Init(rType string, conf routers.Config) routers.Route {
	switch rType {
	case "GIN":
		return gin.NewMyRouter(conf)
	case "FIBER":
		return fiber.NewFiberRouter(conf)
	default:
		return gin.NewMyRouter(conf)
	}
}
