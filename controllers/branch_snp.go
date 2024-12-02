package controller

import (
	"net/http"
	"stock-inventory/config"
	"stock-inventory/models"

	"github.com/gin-gonic/gin"
)


func BranchStockPost(c *gin.Context) {
	var Branchs *models.BranchStockPrice
	err := c.ShouldBind(&Branchs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}

	res := config.DB.Create(Branchs)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Branch cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Branch Created",
	})
} 

func BranchStockShow(c *gin.Context) {
	var Branchs models.BranchStockPrice
	id := c.Param("id")
	res := config.DB.Find(&Branchs, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Branch not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Branchs,
	})
}

func BranchStockUpdate(c *gin.Context) {
	var Branchs models.BranchStockPrice
	id := c.Param("id")
	err := c.ShouldBind(&Branchs)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateBranch models.BranchStockPrice
	res := config.DB.Model(&UpdateBranch).Where("id = ?", id).Updates(Branchs)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Branch not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Branch Updated",
	})
}

func BranchStockDelete(c *gin.Context) {
	var Branchs models.BranchStockPrice
	id := c.Param("id")
	res := config.DB.Find(&Branchs, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Branch not found",
		})
		return
	}
	config.DB.Delete(&Branchs)
	c.JSON(http.StatusOK, gin.H{
		"message": "Branch deleted",
	})
}

