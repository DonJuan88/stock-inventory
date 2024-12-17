package models

import "gorm.io/gorm"

type Branch struct {
	gorm.Model
	BranchCode         string `json:"branch_code"`
	BranchName         string `json:"branch_namee"`
	BranchAddress      string `json:"branch_address"`
	ContactPerson      string `json:"contact_person"`
	ContactPersonPhone string `json:"cp_phone"`
	Phone              string `json:"phone"`
	Active	bool	`json:"active"  gorm:"default:true"`
}
