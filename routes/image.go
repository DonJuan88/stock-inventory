package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func  ImageRoutes(rg *gin.RouterGroup) {
	rg.GET("/image", controller.ImageIndex)
	rg.POST("/image", controller.ImagePost)
	rg.PUT("/image/:id", controller.ImageUpdate)
	rg.DELETE("/image/:id", controller.ImageDelete)

}
