package controller

import (
	"errors"
	"net/http"
	"stock-inventory/config"
	"stock-inventory/models"

	"github.com/gin-gonic/gin"
)

func OrderIndex(c *gin.Context) {
	var Order []models.Order

	res := config.DB.Find(&Order)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Order,
	})
}

func OrderPost(c *gin.Context) {
	var Order *models.Order
	err := c.ShouldBind(&Order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}

	res := config.DB.Create(Order)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Order cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Order Created",
	})
}

func OrderShow(c *gin.Context) {
	var Order models.Order
	id := c.Param("id")
	res := config.DB.Find(&Order, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Order not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Order,
	})
}

func OrderUpdate(c *gin.Context) {
	var Order models.Order
	id := c.Param("id")
	err := c.ShouldBind(&Order)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateOrder models.Order
	res := config.DB.Model(&UpdateOrder).Where("id = ?", id).Updates(Order)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Order not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Order Updated",
	})
}

func OrderDelete(c *gin.Context) {
	var Order models.Order
	id := c.Param("id")
	res := config.DB.Find(&Order, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Order not found",
		})
		return
	}
	config.DB.Delete(&Order)
	c.JSON(http.StatusOK, gin.H{
		"message": "Order deleted",
	})
}
