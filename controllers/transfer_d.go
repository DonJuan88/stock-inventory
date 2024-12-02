package controller

import (
	"errors"
	"net/http"
	"stock-inventory/config"
	"stock-inventory/models"

	"github.com/gin-gonic/gin"
)

func TransferDetailIndex(c *gin.Context) {
	var TransferDetail []models.TransferDetail

	res := config.DB.Find(&TransferDetail)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": TransferDetail,
	})
}

func TransferDetailPost(c *gin.Context) {
	var TransferDetail *models.TransferDetail
	err := c.ShouldBind(&TransferDetail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}

	res := config.DB.Create(TransferDetail)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "TransferDetail cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "TransferDetail Created",
	})
}

func TransferDetailShow(c *gin.Context) {
	var TransferDetail models.TransferDetail
	id := c.Param("id")
	res := config.DB.Find(&TransferDetail, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "TransferDetail not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": TransferDetail,
	})
}

func TransferDetailUpdate(c *gin.Context) {
	var TransferDetail models.TransferDetail
	id := c.Param("id")
	err := c.ShouldBind(&TransferDetail)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateTransferDetail models.TransferDetail
	res := config.DB.Model(&UpdateTransferDetail).Where("id = ?", id).Updates(TransferDetail)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "TransferDetail not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "TransferDetail Updated",
	})
}

func TransferDetailDelete(c *gin.Context) {
	var TransferDetail models.TransferDetail
	id := c.Param("id")
	res := config.DB.Find(&TransferDetail, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "TransferDetail not found",
		})
		return
	}
	config.DB.Delete(&TransferDetail)
	c.JSON(http.StatusOK, gin.H{
		"message": "TransferDetail deleted",
	})
}
