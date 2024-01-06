package database

import (
	"user-management/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectDB is a function that connects to the database

func ConnectDB(dbName string) *gorm.DB {

	db, err := gorm.Open(sqlite.Open(dbName+".db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{})
	return db
}
