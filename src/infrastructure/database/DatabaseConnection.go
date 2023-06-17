package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect() *mongo.Client {
	var user string
	if user = os.Getenv("MONGODB_USER"); user == "" {
		user = "root"
	}
	var pass string
	if pass = os.Getenv("MONGODB_PASS"); pass == "" {
		pass = "root"
	}
	var host string
	if host = os.Getenv("MONGODB_HOST"); host == "" {
		host = "localhost"
	}
	var port string
	if port = os.Getenv("MONGODB_PORT"); port == "" {
		port = "27017"
	}
	uri := "mongodb://" + user + ":" + pass + "@" + host + ":" + port
	log.SetPrefix("[MongoConnect] ")
	log.Print("Connecting to " + uri)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
