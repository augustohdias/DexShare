package adapter

import (
	"dexshare/src/core/entity"
	"dexshare/src/port"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserAdapter struct {
	UserService port.UserServicePort
}

func (u *UserAdapter) GetUser(c *gin.Context) {
	id := c.Param("uid")
	user, err := u.UserService.Read(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	type Response struct {
		ID        string   `json:"id"`
		Name      string   `json:"name"`
		Pokemons  []string `json:"pokemons"`
		Followers []string `json:"followers"`
		Following []string `json:"following"`
	}
	pokemons := user.Pokemons
	if pokemons == nil {
		pokemons = []string{}
	}
	followers := user.Followers
	if followers == nil {
		followers = []string{}
	}
	following := user.Following
	if following == nil {
		following = []string{}
	}
	c.JSON(http.StatusOK, Response{
		ID:        user.ID,
		Name:      user.Name,
		Pokemons:  pokemons,
		Followers: followers,
		Following: following,
	})
	c.Next()
}

func (u *UserAdapter) UploadSaveFile(c *gin.Context) {
	type RequestBody struct {
		File string `json:"file" validate:"required"`
	}
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	id := c.Param("uid")
	_, err := u.UserService.UploadSaveFile(id, body.File)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusAccepted)
	c.Next()
}

func (u *UserAdapter) CreateUser(c *gin.Context) {
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
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Mandatory fields are missing.", "missingFields": missingFields})
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	id, err := u.UserService.Create(user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, Response{ID: id})
	c.Next()
}
