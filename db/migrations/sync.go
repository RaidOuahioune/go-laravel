package migrations

import (
	"demo.com/hello/db"
	"demo.com/hello/models"
)

func SyncTableSchemas() {

	var database = &db.Database{}
	var db = database.GetInstance()

	for _, table := range models.Tables {
		db.Migrator().AutoMigrate(table)
	}

}
