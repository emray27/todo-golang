package app

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var mongoError error
var todoCollection *mongo.Collection

func MongoConnectInit() {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		fmt.Println("URI was not loaded properly")
	}

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		mongoError = err
		fmt.Println("Could not connect to mongoDB")
		return
	}

	// Test connection
	error := client.Ping(context.TODO(), nil)
	if error != nil {
		mongoError = error
		fmt.Println("Connect ping failed!")
		return
	}

	mongoClient = client

	fmt.Println("Connected to mongo client")
}

func GetMongoClient() (*mongo.Client, error) {
	if mongoClient == nil {
		MongoConnectInit()
	}
	return mongoClient, mongoError
}

func GetTodosCollection() (*mongo.Collection, error) {
	mongoClient, error := GetMongoClient()
	if error != nil {
		mongoError = error
	} else {
		todoCollection = mongoClient.Database("todo_project").Collection("todos")
	}
	return todoCollection, mongoError
}
