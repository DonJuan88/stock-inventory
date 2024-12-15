package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func BrandRoutes(rg *gin.RouterGroup) {
	rg.GET("/brands/:id", controller.BrandShow)
	rg.GET("/brands", controller.BrandIndex)
	rg.POST("/brands", controller.BrandPost)
	rg.PUT("/brands/:id", controller.BrandUpdate)
	rg.DELETE("/brands/:id", controller.BrandDelete)

}
