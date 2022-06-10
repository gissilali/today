package repositories

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() {
	db := CurrentDB()
	db.AutoMigrate(&Task{})
}

func CurrentDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
