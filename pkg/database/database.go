package database

import (
	"fmt"
	"go-api-boilerplate/modules/user/user_models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectToDatabase establishes a connection to the SQLite database.
func ConnectToDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.AutoMigrate(&user_models.User{})
	if err != nil {
		return nil, fmt.Errorf("failed to migragte database schema: %w", err)
	}

	fmt.Println("Connect to Database Successfully !")

	return db, nil
}
