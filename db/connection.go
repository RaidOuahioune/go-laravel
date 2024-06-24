package db

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DATABASE_CONFIG = gin.H{
	"max_idle_conns":    10,
	"max_open_conns":    100,
	"conn_max_lifetime": time.Hour,
}

type Database struct {
}

func (m *Database) GetInstance() *gorm.DB {

	dsn := "host=postgres user=root password=raid2019rr dbname=sentry port=5432 sslmode=disable"
	var db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	sqlDB, _ := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(DATABASE_CONFIG["max_idle_conns"].(int))

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(DATABASE_CONFIG["max_open_conns"].(int))

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
