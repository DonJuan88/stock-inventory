package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func PurchaseDRoutes(rg *gin.RouterGroup) {
	rg.GET("/purchasedetails/:id", controller.PurchaseDShow)
	rg.GET("/purchasedetails", controller.PurchaseDIndex)
	rg.POST("/purchasedetails", controller.PurchaseDPost)
	rg.PUT("/purchasedetails/:id", controller.PurchaseDUpdate)
	rg.DELETE("/purchasedetails/:id", controller.PurchaseDDelete)

}
