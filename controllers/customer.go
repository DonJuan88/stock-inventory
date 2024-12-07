package controller

import (
	"errors"
	"net/http"
	"stock-inventory/config"
	"stock-inventory/models"

	"github.com/gin-gonic/gin"
)

func CustomerIndex(c *gin.Context) {
	var Customer []models.Customer
	res := config.DB.Find(&Customer)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Customer,
	})
}

func CustomerPost(c *gin.Context) {
	var Customer *models.Customer
	err := c.ShouldBind(&Customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	res := config.DB.Create(Customer)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Customer cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Customer,
	})
}

func CustomerShow(c *gin.Context) {
	var Customer models.Customer
	id := c.Param("id")
	res := config.DB.Find(&Customer, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Customer not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Customer,
	})
}

func CustomerUpdate(c *gin.Context) {
	var Customer models.Customer
	id := c.Param("id")
	err := c.ShouldBind(&Customer)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateCustomer models.Customer
	res := config.DB.Model(&UpdateCustomer).Where("id = ?", id).Updates(Customer)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Customer not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Customer,
	})
}

func CustomerDelete(c *gin.Context) {
	var Customer models.Customer
	id := c.Param("id")
	res := config.DB.Find(&Customer, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Customer not found",
		})
		return
	}
	config.DB.Delete(&Customer)
	c.JSON(http.StatusOK, gin.H{
		"message": "Customer deleted",
	})
}
