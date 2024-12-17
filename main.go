package main

import (
	"fmt"
	"stock-inventory/config"
	"stock-inventory/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	config.LoadConfig()
	config.DatabaseConnection()
	//	gin.SetMode(gin.ReleaseMode)
	fmt.Println("Starting Application...")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())
	r.SetTrustedProxies(nil)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.SetTrustedProxies(nil)

	//login

	api := r.Group("/api/v1")
	{
		routes.AccountUserRoutes(api)
		routes.BranchRoutes(api)
		routes.BrandRoutes(api)
		routes.CategoryRoutes(api)
		routes.CompanyRoutes(api)
		routes.CustomerRoutes(api)
		routes.ImageRoutes(api)
		routes.MessageRoutes(api)
		routes.NotificationRoutes(api)
		routes.SaleRoutes(api)
		routes.SaleDRoutes(api)
		routes.ProductRoutes(api)
		routes.PurchaseDRoutes(api)
		routes.PurchaseRoutes(api)
		routes.SupplierRoutes(api)
		routes.TransferDRoutes(api)
		routes.TransferRoutes(api)
	}

	//execute
	r.ForwardedByClientIP = true
	r.Run()

}
