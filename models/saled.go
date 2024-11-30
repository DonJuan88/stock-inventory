package models

import (
	"time"

	"gorm.io/gorm"
)

type SaleDetail struct {
	gorm.Model
	SaleNo    string    `json:"sale_no"`
	SaleDate  time.Time `json:"sale_date"`
	ItemCode  string    `json:"code"`
	Qty       int64     `json:"qty"`
	BasePrice int64     `json:"baseprice"`
	Discount  int64     `json:"discount"`
	SalePrice int64     `json:"saleprice"`
}
