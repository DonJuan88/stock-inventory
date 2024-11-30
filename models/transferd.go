package models

import "gorm.io/gorm"

type TransferDetail struct {
	gorm.Model
	TransferNo string `json:"transfer_no"`
	ItemCode   string `json:"code"`
	Qty        int64  `json:"qty"`
}
