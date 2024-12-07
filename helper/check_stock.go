package helper

import (
	"stock-inventory/models"

	"gorm.io/gorm"
)

func CheckStock(db *gorm.DB, locstion string) (*models.BranchStockPrice, error) {
	var checkStock models.BranchStockPrice
	result := db.First(&checkStock, "branch_code = ?", locstion)

	if result.Error != nil {
		return nil, result.Error
	}

	return &checkStock, nil
}
