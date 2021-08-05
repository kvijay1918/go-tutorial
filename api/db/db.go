package db

import (
	"fmt"
	"os"
	"vijay/project-1/api/constants"
	"vijay/project-1/api/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {

	var err error

	db, err = gorm.Open(postgres.Open(constants.DbURL), &gorm.Config{})

	if err != nil {
		fmt.Println("Error in connecting to DB : ", err)
		os.Exit(1)
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println("Unable to creste user table", err)
		os.Exit(1)
	}

}
