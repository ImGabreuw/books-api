package database

import (
	"books-api/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func StartDB() {
	dsn := "host=localhost dbname=books sslmode=disable user=postgres password=root"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error: ", err)
	}

	db = database

	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

func GetDataBase() *gorm.DB {
	return db
}
