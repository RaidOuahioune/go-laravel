package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
}

func (m *Database) GetInstance() *gorm.DB {

	dsn := "host=postgres user=root password=raid2019rr dbname=sentry port=5432 sslmode=disable"
	var db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db
}
