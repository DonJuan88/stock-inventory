package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func SaleDRoutes(rg *gin.RouterGroup) {
	rg.GET("/saledetail/:id", controller.SaleDShow)
	rg.GET("/saledetail", controller.SaleDIndex)
	rg.POST("/saledetail", controller.SaleDPost)
	rg.PUT("/saledetail/:id", controller.SaleDUpdate)
	rg.DELETE("/saledetail/:id", controller.SaleDDelete)

}
