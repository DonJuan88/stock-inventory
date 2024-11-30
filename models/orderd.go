package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderDetail struct {
	gorm.Model
	OrderNo   string    `json:"order_no"`
	OrderDate time.Time `json:"order_date"`
	ItemCode  string    `json:"code"`
	Qty       int64     `json:"qty"`
	Price     int64     `json:"price"`
	Discount  int64     `json:"discount"`
	SalePrice int64     `json:"price_after_disc"`
}
