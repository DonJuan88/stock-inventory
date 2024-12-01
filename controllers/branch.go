package controller

import (
	"errors"
	"net/http"
	"stock-inventory/config"
	"stock-inventory/helper"
	"stock-inventory/models"

	"github.com/gin-gonic/gin"
)

func BranchIndex(c *gin.Context) {
	var Branchs []models.Branch

	res := config.DB.Find(&Branchs)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Branchs,
	})
}

func BranchPost(c *gin.Context) {
	var Branchs *models.Branch
	err := c.ShouldBind(&Branchs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}

	// Check if Branch exists
	exists, err := helper.CheckBranchExists(config.DB, Branchs.BranchCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Branch Code already registered"})
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

func BranchShow(c *gin.Context) {
	var Branchs models.Branch
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

func BranchUpdate(c *gin.Context) {
	var Branchs models.Branch
	id := c.Param("id")
	err := c.ShouldBind(&Branchs)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateBranch models.Branch
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

func BranchDelete(c *gin.Context) {
	var Branchs models.Branch
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
