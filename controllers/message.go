package controller

import (
	"errors"
	"net/http"
	"stock-inventory/config"
	"stock-inventory/models"

	"github.com/gin-gonic/gin"
)

func MessageIndex(c *gin.Context) {
	var Message []models.Message

	res := config.DB.Find(&Message)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Message,
	})
}

func MessagePost(c *gin.Context) {
	var Message *models.Message
	err := c.ShouldBind(&Message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}

	res := config.DB.Create(Message)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Message cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Message Created",
	})
}

func MessageShow(c *gin.Context) {
	var Message models.Message
	id := c.Param("id")
	res := config.DB.Find(&Message, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Message not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Message,
	})
}

func MessageUpdate(c *gin.Context) {
	var Message models.Message
	id := c.Param("id")
	err := c.ShouldBind(&Message)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateMessage models.Message
	res := config.DB.Model(&UpdateMessage).Where("id = ?", id).Updates(Message)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Message not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Message Updated",
	})
}

func MessageDelete(c *gin.Context) {
	var Message models.Message
	id := c.Param("id")
	res := config.DB.Find(&Message, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Message not found",
		})
		return
	}
	config.DB.Delete(&Message)
	c.JSON(http.StatusOK, gin.H{
		"message": "Message deleted",
	})
}
