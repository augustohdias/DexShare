package port

import "dexshare/src/core/entity"

type PokemonRepositoryPort interface {
	Save(pokemon entity.PokemonEntity) (string, error)
	Find(pokemonID string) (entity.PokemonEntity, error)
	Delete(pokemonID string)
}
