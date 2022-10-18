package switchrouter

import (
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/ports"
	router "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/routers"
	introute "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/routers/init"
)

/*SwitchRouter ....*/
type SwitchRouter struct {
	router.Route
}

var sr SwitchRouter

/*New return My Router */
func New(b ports.BookSrv) *SwitchRouter {
	r := introute.Init("GIN", router.Config{
		Port: 3000,
	})

	sr.Route = r
	return &sr

}
