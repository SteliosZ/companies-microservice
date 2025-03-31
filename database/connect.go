package database

import (
	"company/microservice/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBInstance struct {
	DB *gorm.DB
}

var tx DBInstance

type ConnectionConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

// ConnectToDB provides a connection to postgresql
// tx.DB is usable throughout
func ConnectToDB() *DBInstance {

	// Initialize Connection Configuration
	connConfig := ConnectionConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		Port:     os.Getenv("POSTGRES_PORT"),
		SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
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

	// Start db connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("db connection failure: %v", err.Error())
	}

	log.Println("Connection to DB successful")

	// automigrate database from models
	err = db.AutoMigrate(
		&model.Company{},
		&model.User{},
	)
	if err != nil {
		log.Fatalf("database migration failure: %v", err.Error())
	}

	log.Println("Database Migrated successfully")

	tx = DBInstance{
		DB: db,
	}

	return &tx
}
