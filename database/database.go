package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectToDatabase establishes a connection to the SQLite database.
func ConnectToDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil

}

// func MigrateDatabase() {
// 	// Migrate the schema
// 	err = db.AutoMigrate(&models.User{})
// 	if err != nil {
// 		log.Fatal("Failed to migrate schema: ", err)
// 	}
// }

// func

//   db, err := server.ConnectToDatabase()
//         if err != nil {
//                 log.Fatal("Failed to connect to database: ", err)
//         }
//         defer db.Close()

//         // Migrate the schema
//         err = db.AutoMigrate(&models.User{})
//         if err != nil {
//                 log.Fatal("Failed to migrate schema: ", err)
//         }

//         // Run the server (example)
//         err = server.RunServer(db)
//         if err != nil {
//                 log.Fatal("Failed to run server: ", err)
//         }

//         fmt.Println("Server started successfully")
