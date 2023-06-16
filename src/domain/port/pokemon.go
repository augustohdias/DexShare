package port

import "dexshare/src/core/entity"

type PokemonPort interface {
	Create() bool
	Comment(id string, comment string)
	Read(id string) entity.Pokemon
	Update(id string) entity.Pokemon
	Delete(id string)
}
