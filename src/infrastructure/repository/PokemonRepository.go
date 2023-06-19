package repository

import (
	"context"
	"dexshare/src/core/entity"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PokemonRepository struct {
	collection *mongo.Collection
}

func NewPokemonRepository() PokemonRepository {
	p := PokemonRepository{}
	if p.collection == nil {
		client := MongoConnect()
		p.collection = client.Database("dexshare").Collection("pokemon")
	}
	return p
}

func (p *PokemonRepository) Save(pokemon entity.PokemonEntity) (string, error) {
	_, err := p.collection.InsertOne(context.Background(), pokemon)
	if err != nil {
		log.SetPrefix("[PokemonRepository] [Save] ")
		log.Println(err)
		return "", err
	}
	return pokemon.ID, nil
}

func (p *PokemonRepository) Delete(pokemonID string) {
	filter := bson.M{"id": pokemonID}
	_, err := p.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
}

func (p *PokemonRepository) Find(id string) (entity.PokemonEntity, error) {
	var pokemon entity.PokemonEntity
	err := p.collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&pokemon)
	if err != nil {
		log.SetPrefix("[PokemonRepository] [Find] ")
		log.Println(err)
		return entity.PokemonEntity{}, err
	}
	return pokemon, nil
}
