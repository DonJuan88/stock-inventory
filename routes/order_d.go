package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func OrderDRoutes(rg *gin.RouterGroup) {
	rg.GET("/orderdetail/:id", controller.OrderDShow)
	rg.GET("/orderdetail", controller.OrderDIndex)
	rg.POST("/orderdetail", controller.OrderDPost)
	rg.PUT("/orderdetail/:id", controller.OrderDUpdate)
	rg.DELETE("/orderdetail/:id", controller.OrderDDelete)

}
