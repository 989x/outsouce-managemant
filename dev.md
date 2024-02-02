create readme.md with code

gofiber-boilerplate

Layer of app
--- 
src
    api
        configs
        handlers
        helpers
        middleware
        models
        routes
    .env
    go.mod
    go.sum
    main.go
.gitignore
...dockerfile
readme.md
test.http

Layer 1
--- 
    src
    .gitignore
    ...dockerfile
    readme.md
    test.http

Layer 2
--- 
    api
    .env
    go.mod
    go.sum
    main.go

Layer 3
--- 
    configs
    handlers
    helpers
    middleware
    models
    routes


------------------------------------------------------------
1
------------------------------------------------------------

We can take MongoInstance to be announced in src/configs/mongo.go?

src/configs/mongo.go

package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "asdasd"
const mongoURL = "mongodb://127.0.0.1:27017/" + dbName

var MongoClient *mongo.Client

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
	log.Println("Connected to MongoDB")
}



src/models/mongo.go

package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoInstance struct {
	Client *mongo.Client
	Ctx    context.Context
	Db     *mongo.Database
}

------------------------------------------------------------
2
------------------------------------------------------------

Want to take the code in this file and separate it in order to create a better structure.

from
routes/routes.go

to
routes/v1/...

package routes

import (
	"osm/handler"
	"osm/models"

	auth_repository "osm/repository/auth_repo"
	staff_repository "osm/repository/staff_repo"
	user_repository "osm/repository/user_repo"

	auth_service "osm/service/auth_service"
	staff_service "osm/service/staff_service"
	user_service "osm/service/user_service"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, mgConn *models.MongoInstance) {

	api := app.Group("/api")
	v1 := api.Group("/v1")

	authRep := auth_repository.NewAuthRepository()
	authSrv := auth_service.NewAuthRepository(authRep)
	authHan := handler.NewAuthService(authSrv)

	auth := v1.Group("/auth")

	auth.Post("/login", authHan.Login)

	userRep := user_repository.NewUserRepository(mgConn)
	userSrv := user_service.NewUserService(userRep)
	userHan := handler.NewUserHandler(userSrv)

	user := v1.Group("/user")

	user.Get("/getall", userHan.GetAll)
	user.Get("/getby/:id", userHan.GetById)

	staffRep := staff_repository.NewStaffRepoSitory(mgConn)
	staffSrv := staff_service.NewStaffService(staffRep)
	staffHan := handler.NewStaffService(staffSrv)

	staff := v1.Group("/staff")

	staff.Get("/dashboard", staffHan.GetDashboard)
	staff.Get("/getall", staffHan.ListStaffs)
	staff.Get("/getby/:id", staffHan.ReadStaff)
	staff.Get("/testGet", func(c *fiber.Ctx) error {
		result, err := staffRep.GetAllStaffJobs()
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"errors" : err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"data": result,
		})
	})

}

------------------------------------------------------------
3
------------------------------------------------------------

api/repositories/user/user.go:12:17: undefined: models.MongoInstance


/configs/mongo.go

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

/repositories/user/user.go

package repositories

import (
	"dododev/api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	mgConn *models.MongoInstance
}

func NewUserRepository(mgConn *models.MongoInstance) UserRepository {
	return userRepository{mgConn: mgConn}
}

func (r userRepository) GetAll() (*mongo.Cursor, error) {
	collection := r.mgConn.Db.Collection("users")

	query := bson.D{{}}
	query_result, err := collection.Find(r.mgConn.Ctx, query)
	if err != nil {
		return nil, err
	}

	return query_result, nil
}

func (r userRepository) GetById(_id primitive.ObjectID) (*mongo.Cursor, error) {
	collection := r.mgConn.Db.Collection("users")

	query := bson.D{{Key: "_id", Value: _id}}
	query_result, err := collection.Find(r.mgConn.Ctx, query)
	if err != nil {
		return nil, err
	}

	return query_result, nil
}
