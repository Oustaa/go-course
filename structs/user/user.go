package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthday string
	createdAt time.Time
}

type Admin struct  {
	email string
	password string
	User
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("all fields are required")
	}
	return &User{
		firstName,
		lastName,
		birthdate,
		time.Now(),
	}, nil
}

func (u *User) GetFullName() string {
	return fmt.Sprintf("%s %s", u.firstName, u.lastName)
}

func (u *User) ClearUserName () {
	u.firstName = ""
	u.lastName = ""
}

func (u *User) PrintUserInfo() {
	fmt.Printf("Hello %s %s, your birthday is %s\n", u.firstName, u.lastName, u.birthday)
}