package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Avatar   string
}

func (u User) FindById(id uint64) (User, error) {
	return User{Username: "Ousta6", Email: "Otailaba98@gmail.com", Password: "123456789", Avatar: "https://localhost:8000/statics/avatar.png"}, nil
}
