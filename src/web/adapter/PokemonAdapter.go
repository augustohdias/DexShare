package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PokemonController struct{}

func (p *PokemonController) getPokemon(c *gin.Context) {
	// TODO
	c.IndentedJSON(http.StatusOK, "hi lorena")
}
