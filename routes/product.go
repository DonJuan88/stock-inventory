package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func  ProductRoutes(rg *gin.RouterGroup) {
	rg.GET("Product/:id", controller.ProductShow)
	rg.GET("Product", controller.ProductIndex)
	rg.POST("Product", controller.ProductPost)
	rg.PUT("Product/:id", controller.ProductUpdate)
	rg.DELETE("Product/:id", controller.ProductDelete)

}
