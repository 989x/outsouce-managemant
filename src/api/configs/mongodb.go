package configs

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoInstance struct {
	Client *mongo.Client
	Ctx    context.Context
	Db     *mongo.Database
}

const dbName = "SODdb"
const mongoURL = "mongodb://127.0.0.1:27017/" + dbName

var MgConn mongoInstance

func InitMongoDb() {
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
}
