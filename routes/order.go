package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func  OrderDoutes(rg *gin.RouterGroup) {
	rg.GET("/order/:id", controller.OrderShow)
	rg.GET("/order", controller.OrderIndex)
	rg.POST("/order", controller.OrderPost)
	rg.PUT("/order/:id", controller.OrderUpdate)
	rg.DELETE("/order/:id", controller.OrderDelete)

}
