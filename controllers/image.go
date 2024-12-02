package controller

import (
	"errors"
	"net/http"
	"stock-inventory/config"
	"stock-inventory/models"

	"github.com/gin-gonic/gin"
)

func ImageIndex(c *gin.Context) {
	var Image []models.ProductImage

	res := config.DB.Find(&Image)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Image,
	})
}

func ImagePost(c *gin.Context) {
	var Image *models.ProductImage
	err := c.ShouldBind(&Image)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}

	res := config.DB.Create(Image)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Image cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Image Created",
	})
}

func Imagehow(c *gin.Context) {
	var Image models.ProductImage
	id := c.Param("id")
	res := config.DB.Find(&Image, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Image not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Image,
	})
}

func ImageUpdate(c *gin.Context) {
	var Image models.ProductImage
	id := c.Param("id")
	err := c.ShouldBind(&Image)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateImage models.ProductImage
	res := config.DB.Model(&UpdateImage).Where("id = ?", id).Updates(Image)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Image not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Image Updated",
	})
}

func ImageDelete(c *gin.Context) {
	var Image models.ProductImage
	id := c.Param("id")
	res := config.DB.Find(&Image, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Image not found",
		})
		return
	}
	config.DB.Delete(&Image)
	c.JSON(http.StatusOK, gin.H{
		"message": "Image deleted",
	})
}
