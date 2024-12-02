package controller

import (
	"net/http"
	"stock-inventory/config"
	"stock-inventory/helper"
	"stock-inventory/models"
	"stock-inventory/utility"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)



func UserPost(c *gin.Context) {
	var User *models.User
	err := c.ShouldBind(&User)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	exists, err := helper.CheckEmailExists(config.DB, User.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "User Code already registered"})
		return
	}

	res := config.DB.Create(User)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User cannot created",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": User,
	})
}

func UserShow(c *gin.Context) {
	var User models.User
	id := c.Param("id")
	res := config.DB.Find(&User, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": User,
	})
}



func UserDelete(c *gin.Context) {
	var User models.User
	id := c.Param("id")
	res := config.DB.Find(&User, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}
	config.DB.Delete(&User)
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
	})
}


func UserUpdatePassword(c *gin.Context) {
	var User models.User
	id := c.Param("id")
	// Cari user berdasarkan ID
	if err := config.DB.Where("id = ?", c.Param("id")).First(&User).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var UpdateUser models.UpdateUser

	if err := c.ShouldBindJSON(&UpdateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(UpdateUser.LastPassword)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	if UpdateUser.Password != UpdateUser.PasswordConfirmation {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password not Match",
		})
		return

	}

	paswordHash, err := utility.HashPassword(UpdateUser.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password cannot be Decrypt",
		})
		return
	}

	//res := config.DB.Model(&UpdateUser).Where("id = ?", id).Updates(User)
	config.DB.Model(&User).Where("id = ?", id).Update("Password", paswordHash)

	c.JSON(http.StatusOK, gin.H{
		"message": "Password Updated",
	})
}


func Logout(c *gin.Context) {
	authHeader, _ := c.Cookie("Author")

	//ojo diganti

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.SetCookie("Author", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "User logged out successfully",
	})
}
