package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name               string `json:"name"`
	ContactPerson      string `json:"cp"`
	ContactPersonPhone string `json:"cp_phone"`
	Email              string `json:"email"`
	Phone              string `json:"phone"`
	Address            string `json:"address"`
	City               string `json:"city"`
	State              string `json:"state"`
	PostalCode         string `json:"postalcode"`
	Country            string `json:"country"`
}
