package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string  `gorm:"not null" validate:"required"`
	Email     string  `gorm:"unique;not null" validate:"required,email"`
	Age       int     `gorm:"not null" validate:"required,gt=0"`
	Password  string  `gorm:"not null" validate:"required,min=6"`
	CompanyID int     `gorm:"default:null"`
	Company   Company `gorm:"foreignKey:CompanyID" json:"-"`
	Todos     []Todo  `json:"-"`
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

type NewTodo struct {
	Text   string `json:"text" validate:"required"`
	UserID int    `json:"userId"`
}
type Todo struct {
	gorm.Model
	Text   string `gorm:"not null" validate:"required"`
	Done   bool   `gorm:"default:false" validate:"required"`
	UserID int    `gorm:"not null" validate:"gt=0,required"`
	User   User   `gorm:"foreignKey:UserID"`
}

var Tables = []interface{}{&User{}, &Company{}, &Location{}, &Todo{}}
var Validate *validator.Validate

func InitValidation() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}
