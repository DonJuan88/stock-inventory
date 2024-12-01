package helper

import (
	"stock-inventory/models"

	"gorm.io/gorm"
)

func CheckBranchExists(db *gorm.DB, code string) (bool, error) {
	var Branch models.Branch
	result := db.Where("Branch_code = ?", code).First(&Branch)

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
