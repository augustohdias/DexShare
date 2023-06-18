package port

import "dexshare/src/core/entity"

type PokemonRepositoryPort interface {
	Save(entity.PokemonEntity) (string, error)
	Find(string) (entity.PokemonEntity, error)
}
