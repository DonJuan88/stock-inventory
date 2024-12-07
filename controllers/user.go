package controller

import (
	"net/http"
	"stock-inventory/config"
	"stock-inventory/helper"
	"stock-inventory/models"
	"stock-inventory/utility"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func UserPost(c *gin.Context) {
	var register *models.Register

	if err := c.ShouldBindJSON(&register); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Complete Your Field",
		})
		return
	}

	if register.Password != register.PasswordConfirmation {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password not Match",
		})
		return

	}

	// Check if email exists
	exists, err := helper.CheckEmailExists(config.DB, register.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	paswordHashEmail, err := helper.HashPassword(register.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password cannot be Decrypt",
		})
		return
	}

		paswordHash, err := helper.HashPassword(register.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password cannot be Decrypt",
		})
		return
	}

	uuid := uuid.New()
	uuidString := uuid.String()

	account := models.User{
		Uuid:      uuidString,
		FirstName: register.FirstName,
		LastName:  register.LastName,
		Email:     paswordHashEmail,
		Password:  paswordHash,
		IsAdmin:   register.IsAdmin,
		Active:    true,
	}

	//fmt.Println(account)

	if err := config.DB.Create(&account).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Create Account",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Account Created Successfully",
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
