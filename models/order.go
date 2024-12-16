package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderNo       string    `json:"order_no"`
	OrderDate     time.Time `json:"order_date"`
	ShippingCost  int64     `json:"shippingprice"`
	Tax1          int64     `json:"tax1"`
	Tax2          int64     `json:"tax2"`
	Total         int64     `json:"total"`
	AccountID     string    `json:"accid"`
	PaymentType   string    `json:"paymenttype"`
	Reference     string    `json:"reference"`
	Notes         string    `json:"notes"`
	PaymentStatus bool      `jaon:"status gorm:default:false"`
}
