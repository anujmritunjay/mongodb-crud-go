package main

import (
	"net/http"

	"github.com/anujmritunjay/golang-mongodb-crud/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Hello from the Go lang server",
		})
	})
	routes.RegisterRoutes(r)
	r.Run()

}
