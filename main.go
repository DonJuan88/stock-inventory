package main

import (
	"fmt"
	"stock-inventory/config"
	"stock-inventory/middleware"
	"stock-inventory/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.DatabaseConnection()
	//	gin.SetMode(gin.ReleaseMode)
	fmt.Println("Starting Application...")
	config.DatabaseConnection()

	fmt.Println("Application are Ready to Use")

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	//r.Use(cors.Default())
	

	r.SetTrustedProxies(nil)
	
	//login
	
	
	api := r.Group("/api/v1/", middleware.CheckAuth)
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
		routes.OrderDRoutes(api)
		routes.OrderDoutes(api)
		routes.ProductRoutes(api)
		routes.SaleDRoutes(api)
		routes.SaleRoutes(api)
		routes.SupplierRoutes(api)
		routes.TransferDRoutes(api)
		routes.TransferRoutes(api)
	}


	//execute
	r.ForwardedByClientIP = true
	r.Run(fmt.Sprintf(":%v", config.ENV.URL_PORT))
}
