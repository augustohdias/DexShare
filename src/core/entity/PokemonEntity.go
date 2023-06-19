package entity

type PokemonEntity struct {
	ID                string `bson:"id"`
	OwnerID           string `bson:"ownerId"`
	Name              string `bson:"name"`
	Level             int    `bson:"level"`
	NationalDexNumber int    `bson:"nationalDexNumber"`
}
