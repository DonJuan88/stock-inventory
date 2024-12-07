package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is uploaded"})
		return
	}

	// Save file to the local file system
	pathmaster := filepath.Dir("C:/PROJECT/FRONTEND/FLUTTER/newpawpaw/datase/")
	filename := filepath.Base(file.Filename)
	filePath := filepath.Join(pathmaster, "images", filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	fmt.Println(filePath)

	code := c.PostForm("code")

	item := models.ProductImage{
		ProductCode: code,
		FileName:    filePath,
	}
	fmt.Println(item)
	if err := config.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save data to database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

func ImageShow(c *gin.Context) {
	var images []models.ProductImage
	code := c.Param("item_code")
	res := config.DB.Find(&images, code)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Image not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": images,
	})
}

func ImageDelete(c *gin.Context) {
	var images models.ProductImage

	id := c.Param("id")
	res := config.DB.Find(&images, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Image not found",
		})
		return
	}

	var result string
	config.DB.Raw("select file_name from item_images Where id= ? ", id).Scan(&result)

	fmt.Println(result)
	if err := os.Remove(result); err != nil {
		log.Fatal(err)
	}

	config.DB.Delete(&images)
	c.JSON(http.StatusOK, gin.H{
		"message": "Image deleted",
	})

}
