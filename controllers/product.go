package controller

import (
	"errors"
	"net/http"
	"stock-inventory/config"
	"stock-inventory/models"

	"github.com/gin-gonic/gin"
)

func ProductIndex(c *gin.Context) {
	var Product []models.Product

	res := config.DB.Find(&Product)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Product,
	})
}

func ProductPost(c *gin.Context) {
	var Product *models.Product
	err := c.ShouldBind(&Product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}

	res := config.DB.Create(Product)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Product cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Product Created",
	})
}

func Producthow(c *gin.Context) {
	var Product models.Product
	id := c.Param("id")
	res := config.DB.Find(&Product, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Product,
	})
}

func ProductUpdate(c *gin.Context) {
	var Product models.Product
	id := c.Param("id")
	err := c.ShouldBind(&Product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateProduct models.Product
	res := config.DB.Model(&UpdateProduct).Where("id = ?", id).Updates(Product)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Product not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Product Updated",
	})
}

func ProductDelete(c *gin.Context) {
	var Product models.Product
	id := c.Param("id")
	res := config.DB.Find(&Product, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
		})
		return
	}
	config.DB.Delete(&Product)
	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted",
	})
}
