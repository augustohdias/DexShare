package main

import (
	"dexshare/src/web/adapter"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	a := adapter.Default()
	r := gin.Default()
	setupRoutes(r, a.Routes(), a.LoginAdapter.Authenticate)
	err := r.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal("Couldn't start the application.")
		return
	}
}

func setupRoutes(router *gin.Engine, routes []adapter.Route, authenticate func(c *gin.Context)) {
	for _, route := range routes {
		authRoutes := router.Group("/")
		authRoutes.Use(authenticate)
		if route.Secured {
			handler := route.Func
			router.Handle(route.Method, route.Path, func(c *gin.Context) {
				handler(c)
			})
			continue
		}
		handler := route.Func
		router.Handle(route.Method, route.Path, func(c *gin.Context) {
			handler(c)
		})
	}
}
