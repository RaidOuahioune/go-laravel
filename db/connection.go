package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
}

func (m *Database) GetInstance() *gorm.DB {
	dsn := "root@tcp(database:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db
}
