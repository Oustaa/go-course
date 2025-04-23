package main

import (
	"log"
	"net/http"
	"os"

	"github.com/oustaa/practice-http/database"
	"github.com/oustaa/practice-http/handlers"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate)

	db := database.DbConnection(logger)

	db.AutoMigrate(&database.User{})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		logger.Fatalf("ListenAndServe: %v\n", err)
	}

	http.HandleFunc("/user", handlers.UserHandler.DisplayUserById)
	http.ListenAndServe(":8000", nil)
}
