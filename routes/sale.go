package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func PurchaseRoutes(rg *gin.RouterGroup) {
	rg.GET("/purchase/:id", controller.PurchaseShow)
	rg.GET("/purchase", controller.PurchaseIndex)
	rg.POST("/purchase", controller.PurchasePost)
	rg.PUT("/purchase/:id", controller.PurchaseUpdate)
	rg.DELETE("/purchase/:id", controller.PurchaseDelete)

}
