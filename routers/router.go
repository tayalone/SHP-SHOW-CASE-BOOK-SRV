package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/core/ports"
)

/*MyRouter is Define Attr of SPI Router in Application*/
type MyRouter struct {
	*gin.Engine
}

var mr MyRouter

/*New return My Router */
func New(b ports.BookSrv) *MyRouter {
	r := gin.Default()
	r.GET("/book/:id", func(c *gin.Context) {
		/*GetIDUri is Get ID From Parmas */
		type getIDUri struct {
			ID uint `uri:"id" binding:"required"`
		}

		var gi getIDUri
		if err := c.ShouldBindUri(&gi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		b, errGetPk := b.GetByID(gi.ID)
		if errGetPk != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": errGetPk.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"book":    b,
		})

	})

	mr.Engine = r
	return &mr
}

/*Start is  a trigger Router Serve API*/
func (mr MyRouter) Start() {

	mr.Engine.Run(":3000")
}
