package repository

import (
	"osm/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StaffRepository interface {
	Create(*models.Staff) (*mongo.InsertOneResult, error)
	GetAll() ([]models.StaffResponse, error)
	GetById(primitive.ObjectID) (*mongo.Cursor, error)
	Update(primitive.ObjectID) error
	Remove(primitive.ObjectID) error
}
