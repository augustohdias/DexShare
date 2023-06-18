package entity

type PokemonEntity struct {
	ID                string `bson:"id"`
	Name              string `bson:"name"`
	OwnerID           string `bson:"ownerId"`
	NationalDexNumber int    `bson:"nationalDexNumber"`
	Level             int    `bson:"level"`
}
