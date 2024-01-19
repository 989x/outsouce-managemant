package repository

import (
	"fmt"
	"log"
	"osm/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type staffRepository struct {
	mgConn *models.MongoInstance
}

func NewStaffRepoSitory(mgConn *models.MongoInstance) StaffRepository {
	return staffRepository{mgConn: mgConn}
}

func (r staffRepository) Create(*models.Staff) (*mongo.InsertOneResult, error) {
	return nil, nil
}

func (r staffRepository) GetAll() ([]models.StaffResponse, error) {
	collection := r.mgConn.Db.Collection("staffs")

	query := bson.D{{"team", "Dev"}}
	query_result, err := collection.Find(r.mgConn.Ctx, query)
	if err != nil {
		return nil, err
	}

	var staffs []models.StaffResponse
	if err := query_result.All(r.mgConn.Ctx, &staffs); err != nil {
		return nil, err
	}

	count, err := collection.CountDocuments(r.mgConn.Ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total documents in the collection: %v\n", count)

	return staffs, nil
}

func (r staffRepository) GetById(primitive.ObjectID) (*mongo.Cursor, error) {
	return nil, nil
}
func (r staffRepository) Update(primitive.ObjectID) error {
	return nil
}
func (r staffRepository) Remove(primitive.ObjectID) error {
	return nil
}
