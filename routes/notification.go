package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func  NotificationRoutes(rg *gin.RouterGroup) {
	rg.GET("/notification", controller.NotificationIndex)
	rg.POST("/notification", controller.NotificationPost)
	rg.PUT("/notification/:id", controller.NotificationUpdate)
	rg.DELETE("/notification/:id", controller.NotificationDelete)

}
