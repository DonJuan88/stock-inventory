package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func  SaleDRoutes(rg *gin.RouterGroup) {
	rg.GET("/saledetails/:id", controller.SaleDShow)
	rg.GET("/saledetails", controller.SaleDIndex)
	rg.POST("/saledetails", controller.SaleDPost)
	rg.PUT("/saledetails/:id", controller.SaleDUpdate)
	rg.DELETE("/saledetails/:id", controller.SaleDDelete)

}
