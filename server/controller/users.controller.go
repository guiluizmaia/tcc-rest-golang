package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	Repository "github.com/guiluizmaia/tcc-rest-golang/server/repository"
	Structs "github.com/guiluizmaia/tcc-rest-golang/server/structs"
)

var users = []Structs.User{}

type UsersController struct {
	router     *gin.Engine
	repository Repository.UsersRepository
}

func (usersController UsersController) NewUsersController(router *gin.Engine) {
	router.GET("/", usersController.get)
	router.POST("/", usersController.post)
	usersController.repository = Repository.UsersRepository{}
}

func (usersController UsersController) get(context *gin.Context) {
	users := usersController.repository.List()
	context.IndentedJSON(http.StatusOK, users)
}

func (usersController UsersController) post(context *gin.Context) {
	var newUser Structs.User

	if err := context.BindJSON(&newUser); err != nil {
		context.IndentedJSON(http.StatusBadRequest, err)
	}
	usersController.repository.Create(newUser)
	context.IndentedJSON(http.StatusCreated, newUser)
}
