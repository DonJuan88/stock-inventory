package helper

import (
	"stock-inventory/models"

	"gorm.io/gorm"
)

func CheckLastOrder(db *gorm.DB, Order, OrderNo string) (*models.Order, error) {
	var lastOrder models.Order
	result := db.Table(Order).Select(OrderNo).Order("created_at DESC").Limit(1).Scan(&lastOrder)

	if result.Error != nil {
		return nil, result.Error
	}

	return &lastOrder, nil
}

func CheckLastPurchase(db *gorm.DB, Purchase, purchase_no string) (*models.Purchase, error) {
	var lastPurchase models.Purchase
	result := db.Table(Purchase).Select(purchase_no).Order("created_at DESC").Limit(1).Scan(&lastPurchase)

	if result.Error != nil {
		return nil, result.Error
	}

	return &lastPurchase, nil
}

func CheckLastTransfer(db *gorm.DB, Transfer, transfer_no string) (*models.Transfer, error) {
	var lastTransfer models.Transfer
	result := db.Table(Transfer).Select(transfer_no).Order("created_at DESC").Limit(1).Scan(&lastTransfer)

	if result.Error != nil {
		return nil, result.Error
	}

	return &lastTransfer, nil
}

/*
func CheckLastTransfer(db *gorm.DB) (*models.Transfer, error) {
	var lastTransfer models.Transfer
	result := db.Order("created_at DESC").First(&lastTransfer)

	if result.Error !=nil{
		return nil, result.Error
	}

	return &lastTransfer, nil
}
*/
