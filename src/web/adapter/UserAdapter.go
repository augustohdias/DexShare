package adapter

import (
	"dexshare/src/core/entity"
	"dexshare/src/port"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	UserService port.UserServicePort
}

func (u *UserController) GetUser(c *gin.Context) {
	type Response struct {
		ID        string   `json:"id"`
		Name      string   `json:"name"`
		Followers []string `json:"followers"`
		Following []string `json:"following"`
	}
	id := c.Param("uid")
	user, err := u.UserService.Read(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, "")
		return
	}
	c.IndentedJSON(http.StatusOK, Response{
		ID:        user.ID,
		Name:      user.Name,
		Followers: user.Followers,
		Following: user.Following,
	})
}

func (u *UserController) CreateUser(c *gin.Context) {
	type Response struct {
		ID string `json:"id"`
	}
	var user entity.UserEntity
	if err := c.ShouldBindJSON(&user); err != nil {
		if e, ok := err.(validator.ValidationErrors); ok {
			var missingFields []string
			for _, fieldErr := range e {
				missingFields = append(missingFields, fieldErr.Field())
			}
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Mandatory fields are missing.", "missingFields": missingFields})
			return
		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := u.UserService.Create(user)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "")
		return
	}
	c.IndentedJSON(http.StatusOK, Response{ID: id})
}
