package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func SaleRoutes(rg *gin.RouterGroup) {
	rg.GET("/sale/:id", controller.SaleShow)
	rg.GET("/sale", controller.SaleIndex)
	rg.POST("/sale", controller.SalePost)
	rg.PUT("/sale/:id", controller.SaleUpdate)
	rg.DELETE("/sale/:id", controller.SaleDelete)

}
