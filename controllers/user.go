package controller

import (
	"net/http"
	"stock-inventory/config"
	"stock-inventory/helper"
	"stock-inventory/models"

	"github.com/gin-gonic/gin"
)



func AccountPost(c *gin.Context) {
	var Account *models.User
	err := c.ShouldBind(&Account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	exists, err := helper.CheckEmailExists(config.DB, Account.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Account Code already registered"})
		return
	}

	res := config.DB.Create(Account)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Account cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Account,
	})
}

func AccountShow(c *gin.Context) {
	var Account models.User
	id := c.Param("id")
	res := config.DB.Find(&Account, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Account not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Account,
	})
}

func AccountUpdate(c *gin.Context) {
	var Account models.User
	id := c.Param("id")
	err := c.ShouldBind(&Account)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateAccount models.User
	res := config.DB.Model(&UpdateAccount).Where("id = ?", id).Updates(Account)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Account not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Account,
	})
} 

func AccountDelete(c *gin.Context) {
	var Account models.User
	id := c.Param("id")
	res := config.DB.Find(&Account, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Account not found",
		})
		return
	}
	config.DB.Delete(&Account)
	c.JSON(http.StatusOK, gin.H{
		"message": "Account deleted",
	})
}
