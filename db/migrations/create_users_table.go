package migrations

import (
	"demo.com/hello/db"
	"demo.com/hello/models"
)

func CreateUsersTable() {

	var database = &db.Database{}
	var db = database.GetInstance()

	db.Migrator().CreateTable(&models.User{})

}
