package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func   BrandRoutes(rg *gin.RouterGroup) {
	rg.GET("/brand/:id", controller. BrandShow)
	rg.GET("/brand", controller. BrandIndex)
	rg.POST("/brand", controller. BrandPost)
	rg.PUT("/brand/:id", controller. BrandUpdate)
	rg.DELETE("/brand/:id", controller. BrandDelete)

}
