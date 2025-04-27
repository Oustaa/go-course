package store

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection(logger *log.Logger) *sql.DB {
	db, err := sql.Open("mysql", "user:Password12@/url_shortener")
	if err != nil {
		logger.Fatalf("dbConnection: %v", err)
	}

	return db
}
