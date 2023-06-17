package repository

import (
	"context"
	"dexshare/src/core/entity"
	"dexshare/src/infrastructure/database"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func (u *UserRepository) Connect() *UserRepository {
	if u.collection == nil {
		client := database.MongoConnect()
		u.collection = client.Database("dexshare").Collection("user")
	}
	return u
}

func (u *UserRepository) Save(user entity.UserEntity) (string, error) {
	u = u.Connect()
	_, err := u.collection.InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}
	return user.ID, nil
}

func (u *UserRepository) Find(id string) (entity.UserEntity, error) {
	u = u.Connect()
	var user entity.UserEntity
	err := u.collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&user)
	if err != nil {
		log.Fatal(err.Error())
		return entity.UserEntity{}, err
	}
	return user, nil
}
