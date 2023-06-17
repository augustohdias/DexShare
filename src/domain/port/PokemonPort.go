package port

import "dexshare/src/core/entity"

type PokemonPort interface {
	Create() bool
	Comment(id string, comment string)
	Read(id string) entity.PokemonEntity
	Update(id string) entity.PokemonEntity
	Delete(id string)
}

type PokemonRepositoryPort interface {
	Save(entity.UserEntity) (string, error)
}
