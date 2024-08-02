package main

import (
	"fmt"

	"example.com/structs/user"
)


func main () {
	firstName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name: ")

	ousta, error := user.New(firstName, lastName, "1990-01-01")
	// ousta := user.User{firstName, lastName, "1990-01-01", time.Now()}

	if error != nil {
		// fmt.Println(error)
		panic(error)
	}

	ousta.PrintUserInfo()
	ousta.ClearUserName()
	ousta.PrintUserInfo()
}


func getUserData(msg string) string {
	var value string
	fmt.Print(msg)
	fmt.Scanln(&value)
	return value
}