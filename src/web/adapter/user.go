package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (u *UserController) getUser(c *gin.Context) {
	// TODO
	c.IndentedJSON(http.StatusOK, "hi lorena")
}

func (u *UserController) createUser(c *gin.Context) {
	// TODO
	c.IndentedJSON(http.StatusOK, "{}")
}
