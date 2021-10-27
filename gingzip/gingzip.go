package gingzip

import (
	"log"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func GinRouter() {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.BestCompression))
	r.GET("/", index)

	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}

func index(c *gin.Context) {
	c.String(http.StatusOK, "index")
}
