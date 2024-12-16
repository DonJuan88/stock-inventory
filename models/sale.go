package models

import (
	"time"

	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	SaleNo        string    `json:"sale_no"`
	SaleDate      time.Time `json:"sale_date"`
	BranchCode    string    `json:"branch"`
	ShippingCost  int64     `json:"shippingprice"`
	Tax1          int64     `json:"tax1"`
	Tax2          int64     `json:"tax2"`
	Total         int64     `json:"total"`
	AccountID     string    `json:"accid"`
	PaymentType   string    `json:"paymenttype"`
	ShipStatus    string    `json:"shippingstatus"`
	Reference     string    `json:"reference"`
	Notes         string    `json:"notes"`
	PaymentStatus bool      `jaon:"status gorm:default:false"`
}
