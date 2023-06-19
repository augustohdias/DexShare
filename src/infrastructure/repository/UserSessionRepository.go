package repository

import (
	"context"
	"dexshare/src/core/entity"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserSessionRepository struct {
	collection *mongo.Collection
}

func NewUserSessionRepository() UserSessionRepository {
	u := UserSessionRepository{}
	if u.collection == nil {
		client := MongoConnect()
		u.collection = client.Database("dexshare").Collection("userSession")
	}
	return u
}

func (u *UserSessionRepository) Save(session entity.UserSessionEntity) (string, error) {
	filter := bson.M{"userId": session.UserID}
	update := bson.M{"$set": session}
	opts := options.Update().SetUpsert(true)
	_, err := u.collection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		log.SetPrefix("[UserSessionRepository] [Save] ")
		log.Println(err)
		return "", err
	}
	return session.Key, nil
}

func (u *UserSessionRepository) Find(userId string) (entity.UserSessionEntity, error) {
	var session entity.UserSessionEntity
	err := u.collection.FindOne(context.Background(), bson.M{"id": userId}).Decode(&session)
	if err != nil {
		log.SetPrefix("[UserSessionRepository] [Find] ")
		log.Println(err)
		return entity.UserSessionEntity{}, err
	}
	return session, nil
}
