package schema

import (
	"dexshare/src/core/entity"
)

type GameType int

const (
	FireRedLeafGreen GameType = iota
	RubySapphire
	Emerald
)

func (d GameType) String() string {
	return [...]string{"FireRedLeafGreen", "RubySapphire", "Emerald"}[d]
}

type TrainerInfoSection struct {
	TrainerName string
	TrainerID   int
	Game        GameType
}

type TeamSection struct {
	Size     int
	Pokemons []entity.PokemonEntity
}

type PCSection struct {
	PCBox []PCBufferSection
}

type PCBufferSection struct {
	Pokemons []entity.PokemonEntity
}

type SaveData struct {
	TrainerInfo TrainerInfoSection
	Team        TeamSection
	PC          PCSection
	SaveCount   int
}
