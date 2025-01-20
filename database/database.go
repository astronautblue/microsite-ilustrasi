package database

import (
	"fmt"
	"log"

	"github.com/microsite-ilustrasi/config"
	"github.com/microsite-ilustrasi/models"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDatabase() {
	config.LoadConfig()
	config := config.AppConfig
	fmt.Println("config", config)
	dsn := "sqlserver://" + config.DB_USER + ":" + config.DB_PASSWORD + "@" +
		config.DB_HOST + ":" + config.DB_PORT + "?database=" + config.DB_NAME +
		"&encrypt=disable&trustServerCertificate=true"

	// Set GORM configuration
	gormConfig := &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	var err error
	DB, err = gorm.Open(sqlserver.Open(dsn), gormConfig)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate the schema with error handling
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Printf("Error during auto migration: %v", err)
		// You might want to handle this error differently depending on your needs
	}

	log.Println("Database connected and migrated successfully")
}
