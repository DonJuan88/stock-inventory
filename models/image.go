package models

import "gorm.io/gorm"

type ProductImage struct {
	gorm.Model
	ProductCode string `json:"code"`
	FileName    string `json:"image"`
}

type ImageUploads struct {
	ProductCode string `json:"code"`
	FileName    string `json:"image"`
}
