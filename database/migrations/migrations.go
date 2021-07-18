package migrations

import (
	"books-api/model"
	"gorm.io/gorm"
	"log"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(model.Book{})

	if err != nil {
		log.Fatal("Error: ", err)
	}
}
