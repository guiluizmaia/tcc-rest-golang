package repository

import (
	Structs "github.com/guiluizmaia/tcc-rest-golang/server/structs"
)

var users = []Structs.User{}

type UsersRepository struct{}

func (usersRepository UsersRepository) List() []Structs.User {
	return users
}

func (usersRepository UsersRepository) Create(user Structs.User) Structs.User {
	users = append(users, user)

	return user
}
