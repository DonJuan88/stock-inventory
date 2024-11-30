package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Uuid          string `json:"uid"`
	Name          string `json:"name"`
	ContactPerson string `json:"cp"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	City          string `json:"city"`
	State         string `json:"state"`
	PostalCode    string `json:"postalcode"`
	Country       string `json:"country"`
}
