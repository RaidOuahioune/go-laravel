package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string  `gorm:"not null"`
	Email     string  `gorm:"unique;not null"`
	Age       int     `gorm:"not null"`
	Password  string  `gorm:"not null"`
	CompanyID int     `gorm:"default:null"`
	Company   Company `gorm:"foreignKey:CompanyID"`
}
type Company struct {
	gorm.Model
	Name      string `gorm:"not null"`
	Users     []User
	Locations []Location
}

type Location struct {
	gorm.Model
	Address   string  `gorm:"not null"`
	CompanyID int     `gorm:"default:null"`
	Company   Company `gorm:"foreignKey:CompanyID"`
	Lang      float64
	Lat       float64
}

var Tables = []interface{}{&User{}, &Company{}, &Location{}}
