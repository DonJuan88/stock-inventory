package models

import "gorm.io/gorm"

type BranchStockPrice struct {
	gorm.Model
	BranchCode  string `json:"branch_code"`
	ProductCode string `json:"code"`
	Barcode1    string `json:"barcode1"`
	Barcode2    string `json:"barcode2"`
	BaasePrice       int64  `json:"price"`
	SalePrice int64 `json:"saleprice"`
	Stock       int64  `json:"stock"`
	MinStock int64 `json:"min_stock"`
}
