package models

import "gorm.io/gorm"

type Transfer struct {
	gorm.Model
	TransferNo    string `json:"transfer_no"`
	BranchOrigin  string `json:"branch_origin"`
	BranchDestiny string `json:"branch_destiny"`
	Reference     string `json:"reference"`
	Notes         string `json:"notes"`
	UserId        string `json:"user_id"`
	Cost          int64  `json:"cost"`
}
