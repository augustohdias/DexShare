package entity

type PokemonEntity struct {
	ID                string `bson:"id" validate:"omitempty"`
	Name              string `bson:"name"`
	Level             int    `bson:"level"`
	NationalDexNumber int    `bson:"nationalDexNumber"`
}
