package db

import (
	"fmt"
	"stockingapi/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Function for handling database connection and auto migration
// Scaffold script, Do Not Edit!
// @param dbHost = database host uri
// @param dbUser = database user
// @param dbPass = database pass
// @param dbName = database name
// @param dbPort = database port
func ConnectDB(dbHost string, dbUser string, dbPass string, dbName string, dbPort string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(models.LoadModels()...)

	return db, nil
}
