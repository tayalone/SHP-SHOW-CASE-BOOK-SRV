package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/repos"
	BookRepo "github.com/tayalone/SHP-SHOW-CASE-BOOK-SRV/repos/book"
)

func main() {
	db := repos.New()
	// ------- Make Repository
	bookRepo := BookRepo.New(db)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

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
		b, errGetPk := bookRepo.GetByPk(gi.ID)
		if errGetPk != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": errGetPk.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"book":    b,
		})

	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
