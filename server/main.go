package main

import (
	Controller "github.com/guiluizmaia/tcc-rest-golang/server/controller"
	Infra "github.com/guiluizmaia/tcc-rest-golang/server/infra"
)

func main() {
	router := Infra.InitRest()
	usersController := Controller.UsersController{}
	usersController.NewUsersController(router)
	router.Run("server:8080")
}
