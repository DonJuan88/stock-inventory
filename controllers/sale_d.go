package controller

import (
	"errors"
	"net/http"
	"stock-inventory/config"
	"stock-inventory/models"

	"github.com/gin-gonic/gin"
)

func SaleDIndex(c *gin.Context) {
	var SaleD []models.SaleDetail

	res := config.DB.Find(&SaleD)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": SaleD,
	})
}

func SaleDPost(c *gin.Context) {
	var SaleD *models.SaleDetail
	err := c.ShouldBind(&SaleD)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}

	res := config.DB.Create(SaleD)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "SaleD cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "SaleD Created",
	})
}

func SaleDhow(c *gin.Context) {
	var SaleD models.SaleDetail
	id := c.Param("id")
	res := config.DB.Find(&SaleD, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "SaleD not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": SaleD,
	})
}

func SaleDUpdate(c *gin.Context) {
	var SaleD models.SaleDetail
	id := c.Param("id")
	err := c.ShouldBind(&SaleD)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateSaleD models.SaleDetail
	res := config.DB.Model(&UpdateSaleD).Where("id = ?", id).Updates(SaleD)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "SaleD not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "SaleD Updated",
	})
}

func SaleDDelete(c *gin.Context) {
	var SaleD models.SaleDetail
	id := c.Param("id")
	res := config.DB.Find(&SaleD, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "SaleD not found",
		})
		return
	}
	config.DB.Delete(&SaleD)
	c.JSON(http.StatusOK, gin.H{
		"message": "SaleD deleted",
	})
}
