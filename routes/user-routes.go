package routes

import (
	"github.com/anujmritunjay/golang-mongodb-crud/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/users", controllers.GetAllUsers)
}
