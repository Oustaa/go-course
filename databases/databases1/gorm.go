package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
}

func main() {
	connection_string := "user=ousta password=Password12 host=127.0.0.1 port=5433 dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(connection_string), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	// u := User{FirstName: "Oussama", LastName: "Tailba", Email: "otailaba98@gmail.com"}
	// db.Create(&u)

	// db.Create(&User{FirstName: "John", LastName: "Doe", Email: "john.doe@gmail.com"})
	// db.Create(&User{FirstName: "James", LastName: "Smith", Email: "james.smith@gmail.com"})

	var user User
	db.First(&user, User{FirstName: "Oussama"})

	fmt.Printf("%v\n", user)

	fmt.Println("###################### USERS ######################")
	var users []User
	// db.Find(&users, User{FirstName: "Oussama"})
	db.Find(&users, "id > ? AND first_name = ? ", 1, "Oussama")
	fmt.Printf("Found %d users\n", len(users))
	fmt.Printf("%v\n", users)
}
