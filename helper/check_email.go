package helper

import (
	"stock-inventory/models"

	"gorm.io/gorm"
)

func CheckEmailExists(db *gorm.DB, email string) (bool, error) {
	var account models.User
	result := db.Where("email = ?", email).First(&account)

	// If record found, return true
	if result.RowsAffected > 0 {
		return true, nil
	}

	// If no record found, return false
	if result.Error == gorm.ErrRecordNotFound {
		return false, nil
	}

	// For other database errors, return false and error
	return false, result.Error
}
