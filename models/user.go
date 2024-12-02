package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Uuid      string `json:"uid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"admin" gorm:"default:false"`
	Active    bool   `json:"active"`
}
