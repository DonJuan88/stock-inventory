package helper

import (
	"stock-inventory/models"

	"gorm.io/gorm"
)

func CheckSupplierExists(db *gorm.DB,uuid string) (bool, error) {
	var supplier models.Supplier
	result := db.Where("uuid = ?", uuid).First(&supplier)

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
