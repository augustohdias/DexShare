package adapter

import (
	"dexshare/src/core/service"

	"github.com/gin-gonic/gin"
)

type Adapter struct {
	UserAdapter    UserAdapter
	LoginAdapter   LoginAdapter
	PokemonAdapter PokemonAdapter
}

func Default() Adapter {
	userService := service.DefaultUserService()
	loginService := service.DefaultLoginService()
	return Adapter{
		UserAdapter:    UserAdapter{UserService: &userService},
		LoginAdapter:   LoginAdapter{LoginPort: &loginService},
		PokemonAdapter: PokemonAdapter{},
	}
}

type Route struct {
	Method  string
	Path    string
	Func    func(*gin.Context)
	Secured bool
}

func (a *Adapter) Routes() []Route {
	return []Route{
		{Path: "/user", Func: a.UserAdapter.CreateUser, Method: "POST", Secured: false},
		{Path: "/user/:uid", Func: a.UserAdapter.UploadSaveFile, Method: "PATCH", Secured: true},
		{Path: "/user/:uid", Func: a.UserAdapter.GetUser, Method: "GET", Secured: false},
		{Path: "/login", Func: a.LoginAdapter.Login, Method: "POST", Secured: false},
	}
}
