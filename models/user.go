package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	id   int
	Name string
	Age  int
}
