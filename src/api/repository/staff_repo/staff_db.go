package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"osm/api/models"

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

func (r staffRepository) GetDashBoard() ([]models.StaffDashBoard, error) {
	collection := r.mgConn.Db.Collection("staff_jobs")

	result_pipeline := DashboardPipeline()

	query_result, err := collection.Aggregate(r.mgConn.Ctx, result_pipeline)
	if err != nil {
		return nil, err
	}

	var staff_jos []models.StaffDashBoard
	if err := query_result.All(r.mgConn.Ctx, &staff_jos); err != nil {
		return nil, err
	}
	return staff_jos, nil
}

func (r staffRepository) Create(*models.Staff) (*mongo.InsertOneResult, error) {
	return nil, nil
}

func (r staffRepository) CallUserExists(buffer *bytes.Buffer) ([]models.StaffExists, error) {
	request, err := http.NewRequest("POST", "https://inet-ra.inet.co.th/api/vMonk/externalApi/searchAllUser_v2", nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	var result_exists []models.StaffExists
	err = json.Unmarshal(result, &result_exists)

	return result_exists, nil
}

func (r staffRepository) GetAll() ([]models.StaffResponse, error) {
	collection := r.mgConn.Db.Collection("staffs")

	query := bson.D{{}}
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

func (r staffRepository) GetAllStaffJobs() ([]models.StaffJobsFind, error) {
	collection := r.mgConn.Db.Collection("staff_jobs")

	query := bson.D{{}}
	query_result, err := collection.Find(r.mgConn.Ctx, query)
	if err != nil {
		return nil, err
	}

	var staffs []models.StaffJobsFind
	if err := query_result.All(r.mgConn.Ctx, &staffs); err != nil {
		return nil, err
	}

	return staffs, nil
}

func (r staffRepository) GetById(_id primitive.ObjectID) (*models.Staff, error) {
	collectoin := r.mgConn.Db.Collection("staffs")

	var staff models.Staff
	query := bson.D{{Key: "_id", Value: _id}}
	if err := collectoin.FindOne(r.mgConn.Ctx, query).Decode(&staff); err != nil {
		return nil, err
	}

	return &staff, nil
}

func (r staffRepository) Update(primitive.ObjectID) error {
	return nil
}

func (r staffRepository) Remove(primitive.ObjectID) error {
	return nil
}
