package routes

import (
	controller "stock-inventory/controllers"

	"github.com/gin-gonic/gin"
)

func   CategoryRoutes(rg *gin.RouterGroup) {
	rg.GET("/category/:id", controller. CategoryShow)
	rg.GET("/category", controller. CategoryIndex)
	rg.POST("/category", controller. CategoryPost)
	rg.PUT("/category/:id", controller. CategoryUpdate)
	rg.DELETE("/category/:id", controller. CategoryDelete)

}
