package migrations

import (
	"gorm.io/gorm"
	"log"
)

func RunMigration(db *gorm.DB) {
	err := db.AutoMigrate()

	if err != nil {
		log.Fatal("Error: ", err)
	}
}
