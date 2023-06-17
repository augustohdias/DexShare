package adapter

import (
	"dexshare/src/core/service"

	"github.com/gin-gonic/gin"
)

type Adapter struct {
	UserController    UserController
	PokemonController PokemonController
}

func Default() Adapter {
	userService := service.DefaultUserService()
	return Adapter{
		UserController:    UserController{UserService: &userService},
		PokemonController: PokemonController{},
	}
}

type Route struct {
	Method string
	Path   string
	Func   func(*gin.Context)
}

func (a *Adapter) Routes() []Route {
	return []Route{
		{Path: "/user", Func: a.UserController.CreateUser, Method: "POST"},
		{Path: "/user/:uid", Func: a.UserController.GetUser, Method: "GET"},
	}
}
