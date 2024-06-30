package migrations

import (
	"log"

	"demo.com/hello/db"
	"demo.com/hello/models"
)

func PopulateDB() {

	var db = (&db.Database{}).GetInstance()
	// Create Companies
	company1 := models.Company{Name: "Tech Innovators"}
	company2 := models.Company{Name: "Health Solutions"}
	db.Create(&company1)
	db.Create(&company2)

	// Create Locations
	location1 := models.Location{Address: "123 Tech St", CompanyID: int(company1.ID), Lang: 34.05, Lat: -118.25}
	location2 := models.Location{Address: "456 Health Ave", CompanyID: int(company2.ID), Lang: 37.77, Lat: -122.42}
	db.Create(&location1)
	db.Create(&location2)

	// Create Users
	user1 := models.User{Name: "Alice", Email: "alice@example.com", Age: 30, Password: "password123", CompanyID: int(company1.ID)}
	user2 := models.User{Name: "Bob", Email: "bob@example.com", Age: 25, Password: "password123", CompanyID: int(company2.ID)}
	user3 := models.User{Name: "Charlie", Email: "charlie@example.com", Age: 35, Password: "password123"}
	db.Create(&user1)
	db.Create(&user2)
	db.Create(&user3)

	// Create Todos
	todo1 := models.Todo{Text: "Complete project", Done: false, UserID: int(user1.ID)}
	todo2 := models.Todo{Text: "Prepare presentation", Done: true, UserID: int(user2.ID)}
	todo3 := models.Todo{Text: "Update resume", Done: false, UserID: int(user3.ID)}
	db.Create(&todo1)
	db.Create(&todo2)
	db.Create(&todo3)

	log.Println("Database populated with sample data")
}
