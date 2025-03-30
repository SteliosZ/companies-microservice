package database

import (
	"company/microservice/config"
	"company/microservice/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBInstance struct {
	DB *gorm.DB
}

var DB DBInstance

type ConnectionConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

func Connect() {

	connConfig := ConnectionConfig{
		Host:     config.Config("POSTGRES_HOST"),
		User:     config.Config("POSTGRES_USER"),
		Password: config.Config("POSTGRES_PASSWORD"),
		DBName:   config.Config("POSTGRES_DB"),
		Port:     config.Config("POSTGRES_PORT"),
		SSLMode:  config.Config("POSTGRES_SSLMODE"),
	}

	// Define Data Source Name
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v",
		connConfig.Host,
		connConfig.User,
		connConfig.Password,
		connConfig.DBName,
		connConfig.Port,
		connConfig.SSLMode,
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

	DB = DBInstance{
		DB: db,
	}
}
