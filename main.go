package main

import (
	"desafio-itau-golang/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.HandleRequest(r)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
