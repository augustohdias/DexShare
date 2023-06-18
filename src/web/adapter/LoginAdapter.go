package adapter

import (
	"dexshare/src/port"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginAdapter struct {
	LoginPort port.LoginPort
}

func (l *LoginAdapter) Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	id := c.Param("uid")
	if !l.LoginPort.Authenticate(id, token) {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}
	c.Next()
}

func (l *LoginAdapter) Login(c *gin.Context) {
	type Request struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	token, err := l.LoginPort.Login(req.Email, req.Password)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}

	type Response struct {
		Token string `json:"token"`
	}
	c.JSON(http.StatusAccepted, Response{
		Token: token,
	})
	c.Next()
}
