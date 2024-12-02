package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)


func AccountUserRoutes( rg *gin.RouterGroup){

	rg.GET("/accounts/:id", controller.UserShow)
	rg.PUT("/accounts/:id", controller.UserUpdatePassword)

}
