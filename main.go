package main

import (
	"github.com/gin-gonic/gin"

	"dexshare/src/web/adapter"
)

func main() {
	adapterRoutes := adapter.Default()
	router := gin.Default()
	getRoutes := append(adapterRoutes.Get())
	postRoutes := append(adapterRoutes.Post())
	setupRoutes(router, "GET", getRoutes)
	setupRoutes(router, "POST", postRoutes)
	router.Run("localhost:8080")
}

func setupRoutes(router *gin.Engine, method string, routes []adapter.HandlerRoute) {
	for _, route := range routes {
		router.Handle(method, route.Path, func(c *gin.Context) {
			route.Func(c)
		})
	}
}
