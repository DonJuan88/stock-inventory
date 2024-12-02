package controller

import (
	"errors"
	"net/http"
	"stock-inventory/config"
	"stock-inventory/models"

	"github.com/gin-gonic/gin"
)

func SupplierIndex(c *gin.Context) {
	var Supplier []models.Supplier

	res := config.DB.Find(&Supplier)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Supplier,
	})
}

func SupplierPost(c *gin.Context) {
	var Supplier *models.Supplier
	err := c.ShouldBind(&Supplier)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}

	res := config.DB.Create(Supplier)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Supplier cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Supplier Created",
	})
}

func Supplierhow(c *gin.Context) {
	var Supplier models.Supplier
	id := c.Param("id")
	res := config.DB.Find(&Supplier, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Supplier not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Supplier,
	})
}

func SupplierUpdate(c *gin.Context) {
	var Supplier models.Supplier
	id := c.Param("id")
	err := c.ShouldBind(&Supplier)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateSupplier models.Supplier
	res := config.DB.Model(&UpdateSupplier).Where("id = ?", id).Updates(Supplier)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Supplier not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Supplier Updated",
	})
}

func SupplierDelete(c *gin.Context) {
	var Supplier models.Supplier
	id := c.Param("id")
	res := config.DB.Find(&Supplier, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Supplier not found",
		})
		return
	}
	config.DB.Delete(&Supplier)
	c.JSON(http.StatusOK, gin.H{
		"message": "Supplier deleted",
	})
}
