package database

import (
	"company/microservice/config"
	"company/microservice/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() {
	// Define Data Source Name
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v",
		config.Config("POSTGRES_HOST"),
		config.Config("POSTGRES_USER"),
		config.Config("POSTGRES_PASSWORD"),
		config.Config("POSTGRES_DB"),
		config.Config("POSTGRES_PORT"),
		config.Config("POSTGRES_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("db connection failure: %v", err.Error())
	}

	log.Println("Connection to DB successful")
	err = db.AutoMigrate(&model.Company{})
	if err != nil {
		log.Fatalf("database migration failure: %v", err.Error())
	}

	log.Println("Database Migrated successfully")
}
