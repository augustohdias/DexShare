package adapter

import (
	"dexshare/src/port"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginService port.LoginPort
}

func (l *LoginController) Login(c *gin.Context) {
	type Request struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	type Response struct {
		Token string `json:"token"`
	}

	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := l.LoginService.Authenticate(req.Email, req.Password)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusUnauthorized)
		return
	}

	c.JSON(http.StatusAccepted, Response{
		Token: token,
	})
}
