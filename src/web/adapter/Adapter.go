package adapter

import (
	"dexshare/src/core/service"
	"dexshare/src/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

type Adapter struct {
	UserAdapter    UserAdapter
	LoginAdapter   LoginAdapter
	PokemonAdapter PokemonAdapter
}

func Default() Adapter {
	userRepository := repository.NewUserRepository()
	pokemonRepository := repository.NewPokemonRepository()
	userSessionRepository := repository.NewUserSessionRepository()

	userService := service.DefaultUserService(&userRepository, &pokemonRepository)
	loginService := service.DefaultLoginService(&userSessionRepository, &userRepository)
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
