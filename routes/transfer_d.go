package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func TransferDRoutes(rg *gin.RouterGroup) {
	rg.GET("/Transferdetail/:id", controller.TransferDetailShow)
	rg.GET("/Transferdetail", controller.TransferDetailIndex)
	rg.POST("/Transferdetail", controller.TransferDetailPost)
	rg.PUT("/Transferdetail/:id", controller.TransferDetailUpdate)
	rg.DELETE("/Transferdetail/:id", controller.TransferDetailDelete)

}
