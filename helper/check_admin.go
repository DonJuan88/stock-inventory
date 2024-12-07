package helper

import (
	"stock-inventory/models"

	"gorm.io/gorm"
)

func CheckStatusUser(db *gorm.DB, id uint) (bool, error) {
	var admin models.User
	result := db.Where("Branch_code = ?", id).First(&admin)

	if result.Error != nil {
		return false, result.Error
	}

	return admin.IsAdmin, nil
}
