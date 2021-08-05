package model

import "gorm.io/gorm"

type User struct {
	gorm.Model //It will update fields like ID, CreatedAt, UpdatedAt fields by default to the DB
	Id         int
	Name       string
	Email      string
	Mobile     string
	Password   string
}
