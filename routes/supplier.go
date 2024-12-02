package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func  SupplierRoutes(rg *gin.RouterGroup) {
	rg.GET("/supplier/:id", controller.SupplierShow)
	rg.GET("/supplier", controller.SupplierIndex)
	rg.POST("/supplier", controller.SupplierPost)
	rg.PUT("/supplier/:id", controller.SupplierUpdate)
	rg.DELETE("/supplier/:id", controller.SupplierDelete)

}
