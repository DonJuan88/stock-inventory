package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(rg *gin.RouterGroup) {
	rg.GET("/categories/:id", controller.CategoryShow)
	rg.GET("/categories", controller.CategoryIndex)
	rg.POST("/categories", controller.CategoryPost)
	rg.PUT("/categories/:id", controller.CategoryUpdate)
	rg.DELETE("/categories/:id", controller.CategoryDelete)

}
