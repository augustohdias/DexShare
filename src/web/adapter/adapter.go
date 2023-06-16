package adapter

import "github.com/gin-gonic/gin"

type HandlerRoute struct {
	Path string
	Func func(*gin.Context)
}

type Adapter struct {
	UserController    UserController
	PokemonController PokemonController
}

func Default() Adapter {
	return Adapter{
		UserController:    UserController{},
		PokemonController: PokemonController{},
	}
}

func (a *Adapter) Get() []HandlerRoute {
	return []HandlerRoute{
		{Path: "/user/:id", Func: a.UserController.getUser},
		{Path: "/pokemon/:id", Func: a.PokemonController.getPokemon},
	}
}

func (a *Adapter) Post() []HandlerRoute {
	return []HandlerRoute{
		{Path: "/user", Func: a.UserController.createUser},
	}
}
