package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func TransferRoutes(rg *gin.RouterGroup) {
	rg.GET("/transferetail/:id", controller.TransferShow)
	rg.GET("/transferetail", controller.TransferIndex)
	rg.POST("/transferetail", controller.TransferPost)
	rg.PUT("/transferetail/:id", controller.TransferUpdate)
	rg.DELETE("/transferetail/:id", controller.TransferDelete)

}
