package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func  BranchRoutes(rg *gin.RouterGroup) {
	rg.GET("/branch/:id", controller.BranchShow)
	rg.GET("/branch", controller.BranchIndex)
	rg.POST("/branch", controller.BranchPost)
	rg.PUT("/branch/:id", controller.BranchUpdate)
	rg.DELETE("/branch/:id", controller.BranchDelete)

}
