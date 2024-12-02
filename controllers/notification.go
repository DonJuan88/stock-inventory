package controller

import (
	"errors"
	"net/http"
	"stock-inventory/config"
	"stock-inventory/models"

	"github.com/gin-gonic/gin"
)

func NotificationIndex(c *gin.Context) {
	var Notification []models.Notification

	res := config.DB.Find(&Notification)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Notification,
	})
}

func NotificationPost(c *gin.Context) {
	var Notification *models.Notification
	err := c.ShouldBind(&Notification)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}

	res := config.DB.Create(Notification)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Notification cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Notification Created",
	})
}

func NotificationShow(c *gin.Context) {
	var Notification models.Notification
	id := c.Param("id")
	res := config.DB.Find(&Notification, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Notification not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Notification,
	})
}

func NotificationUpdate(c *gin.Context) {
	var Notification models.Notification
	id := c.Param("id")
	err := c.ShouldBind(&Notification)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateNotification models.Notification
	res := config.DB.Model(&UpdateNotification).Where("id = ?", id).Updates(Notification)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Notification not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Notification Updated",
	})
}

func NotificationDelete(c *gin.Context) {
	var Notification models.Notification
	id := c.Param("id")
	res := config.DB.Find(&Notification, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Notification not found",
		})
		return
	}
	config.DB.Delete(&Notification)
	c.JSON(http.StatusOK, gin.H{
		"message": "Notification deleted",
	})
}
