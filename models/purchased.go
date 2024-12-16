package models

import (
	"time"

	"gorm.io/gorm"
)

type PurchaseDetail struct {
	gorm.Model
	PurchaseNo    string    `json:"purchase_no"`
	PurchaseDate  time.Time `json:"purchase_date"`
	ItemCode      string    `json:"code"`
	Qty           int64     `json:"qty"`
	BasePrice     int64     `json:"baseprice"`
	Discount      int64     `json:"discount"`
	PurchasePrice int64     `json:"purchaseprice"`
}
