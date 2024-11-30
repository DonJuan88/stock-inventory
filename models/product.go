package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model         //`json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	ProductCode string `json:"code"`
	Barcode1    string `json:"barcode1"`
	Barcode2    string `json:"barcode2"`
	ProductName string `json:"name"`
	Description string `json:"desc"`
	Category    string `json:"category"`
	Brand       string `json:"brand"`
	BasePrice   int64  `json:"baseprice"`
	SalePrice1  int64  `json:"saleprice1"`
	SalePrice2  int64  `json:"saleprice2"`
	SalePrice3  int64  `json:"saleprice3"`
	Unit        string `json:"unit"`
	Active      bool   `json:"active" gorm:"default:true"`
}
