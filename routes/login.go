package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func LoginRoutes(rg *gin.RouterGroup) {

	rg.POST("/login", controller.UserLogin)

	rg.GET("/validate", controller.Validate)
	rg.POST("/logout", controller.Logout)

}
