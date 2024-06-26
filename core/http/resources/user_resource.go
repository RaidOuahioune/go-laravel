package resources

import "demo.com/hello/models"

type UserResource struct {
	models.User
	Company models.Company `json:"company"`
}

func (u *UserResource) TableName() string {
	return "users"
}
