package controller

import (
	"errors"
	"net/http"
	"stock-inventory/config"
	"stock-inventory/models"

	"github.com/gin-gonic/gin"
)

func SaleIndex(c *gin.Context) {
	var Sale []models.Sale

	res := config.DB.Find(&Sale)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Sale,
	})
}

func SalePost(c *gin.Context) {
	var Sale *models.Sale
	err := c.ShouldBind(&Sale)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}

	res := config.DB.Create(Sale)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Sale cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Sale Created",
	})
}

func Salehow(c *gin.Context) {
	var Sale models.Sale
	id := c.Param("id")
	res := config.DB.Find(&Sale, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Sale not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Sale,
	})
}

func SaleUpdate(c *gin.Context) {
	var Sale models.Sale
	id := c.Param("id")
	err := c.ShouldBind(&Sale)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateSale models.Sale
	res := config.DB.Model(&UpdateSale).Where("id = ?", id).Updates(Sale)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Sale not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Sale Updated",
	})
}

func SaleDelete(c *gin.Context) {
	var Sale models.Sale
	id := c.Param("id")
	res := config.DB.Find(&Sale, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Sale not found",
		})
		return
	}
	config.DB.Delete(&Sale)
	c.JSON(http.StatusOK, gin.H{
		"message": "Sale deleted",
	})
}
