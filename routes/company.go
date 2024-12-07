package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func CompanyRoutes(rg *gin.RouterGroup) {
	rg.GET("/company/:id", controller.CompanyShow)
	rg.GET("/company", controller.CompanyIndex)
	rg.POST("/company", controller.CompanyPost)
	rg.PUT("/company/:id", controller.CompanyUpdate)

}
