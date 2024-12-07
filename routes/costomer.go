package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func CustomerRoutes(rg *gin.RouterGroup) {
	rg.GET("/customer/:id", controller.CustomerShow)
	rg.GET("/customer", controller.CustomerIndex)
	rg.POST("/customer", controller.CustomerPost)
	rg.PUT("/customer/:id", controller.CustomerUpdate)
	rg.DELETE("/customer/:id", controller.CustomerDelete)

}
