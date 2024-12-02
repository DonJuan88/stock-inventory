package controller

import (
	"errors"
	"net/http"
	"stock-inventory/config"
	"stock-inventory/models"

	"github.com/gin-gonic/gin"
)

func CompanyIndex(c *gin.Context) {
	var company []models.Company

	res := config.DB.Find(&company)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": company,
	})
}

func CompanyPost(c *gin.Context) {
	var company *models.Company
	err := c.ShouldBind(&company)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}

	res := config.DB.Create(company)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Company cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Company Created",
	})
}

func CompanyShow(c *gin.Context) {
	var company models.Company
	id := c.Param("id")
	res := config.DB.Find(&company, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Company not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": company,
	})
}

func CompanyUpdate(c *gin.Context) {
	var company models.Company
	id := c.Param("id")
	err := c.ShouldBind(&company)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var UpdateCompany models.Company
	res := config.DB.Model(&UpdateCompany).Where("id = ?", id).Updates(company)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Company not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Company Updated",
	})
}

/* func CompanyDelete(c *gin.Context) {
	var company models.Company
	id := c.Param("id")
	res := config.DB.Find(&company, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Company not found",
		})
		return
	}
	config.DB.Delete(&company)
	c.JSON(http.StatusOK, gin.H{
		"message": "Company deleted",
	})
} */
