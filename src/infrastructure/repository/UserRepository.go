package repository

import (
	"context"
	"dexshare/src/core/entity"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository() UserRepository {
	u := UserRepository{}
	if u.collection == nil {
		client := MongoConnect()
		u.collection = client.Database("dexshare").Collection("user")
	}
	return u
}

func (u *UserRepository) Save(user entity.UserEntity) (string, error) {
	_, err := u.collection.InsertOne(context.Background(), user)
	if err != nil {
		log.SetPrefix("[UserRepository] [Save] ")
		log.Println(err)
		return "", err
	}
	return user.ID, nil
}

func (u *UserRepository) UpdatePokemons(user entity.UserEntity) (entity.UserEntity, error) {
	filter := bson.M{"id": user.ID}
	update := bson.M{"$set": user}
	_, err := u.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.SetPrefix("[UserRepository] [Save] ")
		log.Println(err)
		return entity.UserEntity{}, err
	}
	return user, nil
}

func (u *UserRepository) Find(id string) (entity.UserEntity, error) {
	var user entity.UserEntity
	err := u.collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&user)
	if err != nil {
		log.SetPrefix("[UserRepository] [Find] ")
		log.Println(err)
		return entity.UserEntity{}, err
	}
	return user, nil
}

func (u *UserRepository) FindByEmail(email string) (entity.UserEntity, error) {
	var user entity.UserEntity
	err := u.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		log.SetPrefix("[UserRepository] [FindByEmail] ")
		log.Println(err)
		return entity.UserEntity{}, err
	}
	return user, nil
}
