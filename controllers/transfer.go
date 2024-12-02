package controller

import (
	"errors"
	"net/http"
	"stock-inventory/config"
	"stock-inventory/models"

	"github.com/gin-gonic/gin"
)

func TransferIndex(c *gin.Context) {
	var Transfer []models.Transfer

	res := config.DB.Find(&Transfer)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Transfer,
	})
}

func TransferPost(c *gin.Context) {
	var Transfer *models.Transfer
	err := c.ShouldBind(&Transfer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}

	res := config.DB.Create(Transfer)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Transfer cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Transfer Created",
	})
}

func Transferhow(c *gin.Context) {
	var Transfer models.Transfer
	id := c.Param("id")
	res := config.DB.Find(&Transfer, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Transfer not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Transfer,
	})
}

func TransferUpdate(c *gin.Context) {
	var Transfer models.Transfer
	id := c.Param("id")
	err := c.ShouldBind(&Transfer)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateTransfer models.Transfer
	res := config.DB.Model(&UpdateTransfer).Where("id = ?", id).Updates(Transfer)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Transfer not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Transfer Updated",
	})
}

func TransferDelete(c *gin.Context) {
	var Transfer models.Transfer
	id := c.Param("id")
	res := config.DB.Find(&Transfer, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Transfer not found",
		})
		return
	}
	config.DB.Delete(&Transfer)
	c.JSON(http.StatusOK, gin.H{
		"message": "Transfer deleted",
	})
}
