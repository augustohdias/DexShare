package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PokemonController struct{}

func (p *PokemonController) getPokemon(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "hi lorena")
}
