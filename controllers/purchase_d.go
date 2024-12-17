package controller

import (
	"errors"
	"net/http"
	"stock-inventory/config"
	"stock-inventory/models"

	"github.com/gin-gonic/gin"
)

func PurchaseDIndex(c *gin.Context) {
	var PurchaseD []models.PurchaseDetail

	res := config.DB.Find(&PurchaseD)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": PurchaseD,
	})
}

func PurchaseDPost(c *gin.Context) {
	var PurchaseD *models.PurchaseDetail
	err := c.ShouldBind(&PurchaseD)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}

	res := config.DB.Create(PurchaseD)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "PurchaseD cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "PurchaseD Created",
	})
}

func PurchaseDShow(c *gin.Context) {
	var PurchaseD models.PurchaseDetail
	id := c.Param("id")
	res := config.DB.Find(&PurchaseD, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "PurchaseD not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": PurchaseD,
	})
}

func PurchaseDUpdate(c *gin.Context) {
	var PurchaseD models.PurchaseDetail
	id := c.Param("id")
	err := c.ShouldBind(&PurchaseD)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdatePurchaseD models.PurchaseDetail
	res := config.DB.Model(&UpdatePurchaseD).Where("id = ?", id).Updates(PurchaseD)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "PurchaseD not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "PurchaseD Updated",
	})
}

func PurchaseDDelete(c *gin.Context) {
	var PurchaseD models.PurchaseDetail
	id := c.Param("id")
	res := config.DB.Find(&PurchaseD, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "PurchaseD not found",
		})
		return
	}
	config.DB.Delete(&PurchaseD)
	c.JSON(http.StatusOK, gin.H{
		"message": "PurchaseD deleted",
	})
}
