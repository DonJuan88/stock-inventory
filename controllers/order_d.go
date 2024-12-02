package controller

import (
	"errors"
	"net/http"
	"stock-inventory/config"
	"stock-inventory/models"

	"github.com/gin-gonic/gin"
)

func OrderDIndex(c *gin.Context) {
	var OrderD []models.OrderDetail

	res := config.DB.Find(&OrderD)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": OrderD,
	})
}

func OrderDPost(c *gin.Context) {
	var OrderD *models.OrderDetail
	err := c.ShouldBind(&OrderD)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}

	res := config.DB.Create(OrderD)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "OrderD cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OrderD Created",
	})
}

func OrderDShow(c *gin.Context) {
	var OrderD models.OrderDetail
	id := c.Param("id")
	res := config.DB.Find(&OrderD, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "OrderD not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": OrderD,
	})
}

func OrderDUpdate(c *gin.Context) {
	var OrderD models.OrderDetail
	id := c.Param("id")
	err := c.ShouldBind(&OrderD)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateOrderD models.OrderDetail
	res := config.DB.Model(&UpdateOrderD).Where("id = ?", id).Updates(OrderD)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "OrderD not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OrderD Updated",
	})
}

func OrderDDelete(c *gin.Context) {
	var OrderD models.OrderDetail
	id := c.Param("id")
	res := config.DB.Find(&OrderD, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "OrderD not found",
		})
		return
	}
	config.DB.Delete(&OrderD)
	c.JSON(http.StatusOK, gin.H{
		"message": "OrderD deleted",
	})
}
