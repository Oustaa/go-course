package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection(logger *log.Logger) *gorm.DB {
	// db, err := sql.Open("postgres", "user=ousta host=localhost password=Passowrd12 dbname=postgres sslmode=disable")
	dbString := "user=ousta password=Password12 host=localhost port=5433 dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{})
	if err != nil {
		logger.Fatalf("DbConnection: %v\n", err)
	}

	return db
}
