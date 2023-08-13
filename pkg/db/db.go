package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitializeDB(connStr string) *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connStr,
	}), &gorm.Config{})

	if err != nil {
		log.Println("Error connecting to DB", err)
	}

	database, err := db.DB()
	if err != nil {
		log.Println("Initialize DB failed", err)
	}

	database.SetMaxIdleConns(10)
	database.SetMaxOpenConns(50)
	return db
}
