package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func  MessageRoutes(rg *gin.RouterGroup) {
	rg.GET("/message", controller.MessageIndex)
	rg.POST("/message", controller.MessagePost)
	rg.PUT("/message/:id", controller.MessageUpdate)
	rg.DELETE("/message/:id", controller.MessageDelete)

}
