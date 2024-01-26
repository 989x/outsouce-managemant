package database

import (
	"context"
	"fmt"
	"osm/api/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// "context"
// "gomon/models"

const dbName = "SODdb"
const mongoURL = "mongodb://127.0.0.1:27017/" + dbName

func MongoInit() (*models.MongoInstance, error) {
	ctx, cancle := context.WithTimeout(context.Background(), 24*time.Hour)
	_ = cancle
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	if err != nil {
		return nil, err
	}

	mg := models.MongoInstance{
		Client: client,
		Ctx:    ctx,
		Db:     db,
	}

	return &mg, nil
}

var MgConn models.MongoInstance

func MgInit() {
	ctx, cancle := context.WithTimeout(context.Background(), 24*time.Hour)
	_ = cancle
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURL))
	if err != nil {
		panic(err.Error())
	}

	db := client.Database(dbName)
	if err != nil {
		panic(err.Error())
	}

	MgConn.Client = client
	MgConn.Ctx = ctx
	MgConn.Db = db

	fmt.Println("MongoDB Connected")

}
