package controller

import (
	"errors"
	"net/http"
	"stock-inventory/config"
	"stock-inventory/models"

	"github.com/gin-gonic/gin"
)

func PurchaseIndex(c *gin.Context) {
	var Purchase []models.Purchase

	res := config.DB.Find(&Purchase)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Purchase,
	})
}

func PurchasePost(c *gin.Context) {
	var Purchase *models.Purchase
	err := c.ShouldBind(&Purchase)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}

	res := config.DB.Create(Purchase)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Purchase cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Purchase Created",
	})
}

func PurchaseShow(c *gin.Context) {
	var Purchase models.Purchase
	id := c.Param("id")
	res := config.DB.Find(&Purchase, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Purchase not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Purchase,
	})
}

func PurchaseUpdate(c *gin.Context) {
	var Purchase models.Purchase
	id := c.Param("id")
	err := c.ShouldBind(&Purchase)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdatePurchase models.Purchase
	res := config.DB.Model(&UpdatePurchase).Where("id = ?", id).Updates(Purchase)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Purchase not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Purchase Updated",
	})
}

func PurchaseDelete(c *gin.Context) {
	var Purchase models.Purchase
	id := c.Param("id")
	res := config.DB.Find(&Purchase, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Purchase not found",
		})
		return
	}
	config.DB.Delete(&Purchase)
	c.JSON(http.StatusOK, gin.H{
		"message": "Purchase deleted",
	})
}
