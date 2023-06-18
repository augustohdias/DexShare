package adapter

import (
	"dexshare/src/core/service"

	"github.com/gin-gonic/gin"
)

type Adapter struct {
	UserController    UserController
	PokemonController PokemonController
	LoginController   LoginController
}

func Default() Adapter {
	userService := service.DefaultUserService()
	loginService := service.DefaultLoginService()
	return Adapter{
		UserController:    UserController{UserService: &userService},
		LoginController:   LoginController{LoginService: &loginService},
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
		{Path: "/login", Func: a.LoginController.Login, Method: "POST"},
	}
}
