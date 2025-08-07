package repository

import (
	"belajar/entity"
	"errors"
)

// load data, CRUD
func GetUser() (entity.User, error) {

	var users entity.User
	users.Name = "Budi"
	users.Password = "admin"
	users.Username = "admin"

	if false {
		return users, errors.New(("Data tidak ditemuka"))
	}

	return users, nil
}
