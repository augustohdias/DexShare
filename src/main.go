package main

import (
	"github.com/gin-gonic/gin"

	"dexshare/src/web/adapter"
)

func main() {
	adapter := adapter.Default()
	router := gin.Default()
	setupRoutes(router, adapter.Routes())
	router.Run("0.0.0.0:8080")
}

func setupRoutes(router *gin.Engine, routes []adapter.Route) {
	for _, route := range routes {
		handler := route.Func
		router.Handle(route.Method, route.Path, func(c *gin.Context) {
			handler(c)
		})
	}
}
