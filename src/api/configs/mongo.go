package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "SODdb"
const mongoURL = "mongodb://127.0.0.1:27017/" + dbName

var MongoClient *mongo.Client

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var GlobalMongoInstance *MongoInstance // Renamed to avoid conflict

func InitMongoDB() {
	clientOptions := options.Client().ApplyURI(mongoURL)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	MongoClient = client
	GlobalMongoInstance = &MongoInstance{
		Client: client,
		Db:     client.Database(dbName),
	}

	log.Println("Connected to MongoDB")
}
