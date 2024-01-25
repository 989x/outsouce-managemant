package repository

import (
	"osm/api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type staffRepository struct {
	mgConn *models.MongoInstance
}

func NewStaffRepoSitory(mgConn *models.MongoInstance) StaffRepository {
	return staffRepository{mgConn: mgConn}
}

func (r staffRepository) GetStaff() ([]models.StaffDashBoard, error) {
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

func (r staffRepository) SetPrimetiveFilter(available []interface{}, team []interface{}, status []interface{}, center []interface{}) (primitive.D, error) {
	filter_response := bson.D{
		{"$and", bson.A{
			bson.D{{"available", bson.M{"$in": available}}},
			bson.D{{"team", bson.M{"$in": team}}},
			bson.D{{"matchjob", bson.M{"$nin": status}}},
			bson.D{{"canter", bson.M{"$nin": center}}},
		}},
	}
	return filter_response, nil
}

func (r staffRepository) GetPipeLinePrimetive(fillter primitive.D, date_nPlus time.Time) ([]models.StaffDashBoard, error) {
	collection := r.mgConn.Db.Collection("staff_jobs")

	pipeline := []bson.D{
		{
			{"$match", bson.D{
				{"start_jobs_date", bson.D{
					{"$lte", date_nPlus},
				}},
				{"$or", bson.A{
					bson.D{
						{"finish_jobs_date", nil},
					},
					bson.D{
						{"finish_jobs_date", bson.D{
							{"$gte", date_nPlus},
						}},
					},
				}},
			}},
		},
		{
			{"$group", bson.D{
				{"_id", "$user_id"},
				{"matchjob", bson.D{{"$push", "$$ROOT"}}},
			}},
		},
		{
			{"$unwind", "$matchjob"},
		},
		{
			{"$sort", bson.D{
				{"matchjob.available", -1},
				{"matchjob.status", -1},
				{"matchjob.start_jobs_date", -1},
				{"matchjob.updatedAt", -1},
				{"matchjob._id", -1},
			}},
		},
		{
			{"$group", bson.D{
				{"_id", "$_id"},
				{"maxMatchjob", bson.D{{"$first", "$matchjob"}}},
			}},
		},
		{
			{"$lookup", bson.D{
				{"from", "staffs"},
				{"localField", "maxMatchjob.user_id"},
				{"foreignField", "_id"},
				{"as", "maxMatchjob.user"},
			}},
		},
		{
			{"$unwind", "$maxMatchjob.user"},
		},
		{
			{"$project", bson.D{
				{"_id", "$maxMatchjob._id"},
				{"user_id", "$maxMatchjob.user_id"},
				{"start_jobs_date", "$maxMatchjob.start_jobs_date"},
				{"finish_jobs_date", "$maxMatchjob.finish_jobs_date"},
				{"status", "$maxMatchjob.status"},
				{"available", "$maxMatchjob.available"},
				{"outsource", "$maxMatchjob.outsource"},
				{"matchjob", "$maxMatchjob.matchjob"},
				{"address_onsite", "$maxMatchjob.address_onsite"},
				{"status_site", "$maxMatchjob.status_site"},
				{"note", "$maxMatchjob.note"},
				{"createdAt", "$maxMatchjob.createdAt"},
				{"updatedAt", "$maxMatchjob.updatedAt"},
				// from user
				{"id", "$maxMatchjob.user.id"},
				{"fname", "$maxMatchjob.user.fname"},
				{"lname", "$maxMatchjob.user.lname"},
				{"nname", "$maxMatchjob.user.nname"},
				{"start_date", "$maxMatchjob.user.start_date"},
				{"active", "$maxMatchjob.user.active"},
				{"isTransfer", "$maxMatchjob.user.isTransfer"},
				{"last_active_date", "$maxMatchjob.user.last_active_date"},
				{"center", "$maxMatchjob.user.center"},
				{"team", "$maxMatchjob.user.team"},
				{"account_id", "$maxMatchjob.user.account_id"},
			}},
		},
		{
			{"$match", fillter},
		},
		{
			{"$match", bson.D{
				{"start_date", bson.D{{"$lte", date_nPlus}}},
				{"account_id", bson.D{{"$exists", true}}},
				{"account_id", bson.D{{"$ne", nil}}},
				{"$or", bson.A{
					bson.D{{"last_active_date", nil}},
					bson.D{{"last_active_date", bson.D{{"$gt", date_nPlus}}}},
				}},
			}},
		},
	}

	query_result, err := collection.Aggregate(r.mgConn.Ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var getpipeline_response []models.StaffDashBoard

	if err := query_result.All(r.mgConn.Ctx, &getpipeline_response); err != nil {
		return nil, err
	}

	return getpipeline_response, nil
}

func (r staffRepository) GetCountCenterStaff(center string, status string, date_nPlus time.Time) ([]models.StaffCenterStatus, error) {
	collection := r.mgConn.Db.Collection("staff_jobs")

	pipeline := []bson.D{
		{
			{"$match", bson.D{
				{"start_jobs_date", bson.D{
					{"$lte", date_nPlus},
				}},
				{"$or", bson.A{
					bson.D{
						{"finish_jobs_date", nil},
					},
					bson.D{
						{"finish_jobs_date", bson.D{
							{"$gte", date_nPlus},
						}},
					},
				}},
			}},
		},
		{
			{"$group", bson.D{
				{"_id", "$user_id"},
				{"matchjob", bson.D{{"$push", "$$ROOT"}}},
			}},
		},
		{
			{"$unwind", "$matchjob"},
		},
		{
			{"$sort", bson.D{
				{"matchjob.available", -1},
				{"matchjob.status", -1},
				{"matchjob.start_jobs_date", -1},
				{"matchjob.updatedAt", -1},
				{"matchjob._id", -1},
			}},
		},
		{
			{"$group", bson.D{
				{"_id", "$_id"},
				{"maxMatchjob", bson.D{{"$first", "$matchjob"}}},
			}},
		},
		{
			{"$unset", "_id"},
		},
		{
			{"$lookup", bson.D{
				{"from", "staffs"},
				{"localField", "maxMatchjob.user_id"},
				{"foreignField", "_id"},
				{"as", "maxMatchjob.user"},
			}},
		},
		{
			{"$unwind", "$maxMatchjob.user"},
		},
		{
			{"$project", bson.D{
				{"_id", "$maxMatchjob._id"},
				{"user_id", "$maxMatchjob.user_id"},
				{"start_jobs_date", "$maxMatchjob.start_jobs_date"},
				{"finish_jobs_date", "$maxMatchjob.finish_jobs_date"},
				{"status", "$maxMatchjob.status"},
				{"available", "$maxMatchjob.available"},
				{"outsource", "$maxMatchjob.outsource"},
				{"matchjob", "$maxMatchjob.matchjob"},
				{"address_onsite", "$maxMatchjob.address_onsite"},
				{"status_site", "$maxMatchjob.status_site"},
				{"note", "$maxMatchjob.note"},
				{"createdAt", "$maxMatchjob.createdAt"},
				{"updatedAt", "$maxMatchjob.updatedAt"},
				// from user
				{"id", "$maxMatchjob.user.id"},
				{"fname", "$maxMatchjob.user.fname"},
				{"lname", "$maxMatchjob.user.lname"},
				{"nname", "$maxMatchjob.user.nname"},
				{"start_date", "$maxMatchjob.user.start_date"},
				{"active", "$maxMatchjob.user.active"},
				{"isTransfer", "$maxMatchjob.user.isTransfer"},
				{"last_active_date", "$maxMatchjob.user.last_active_date"},
				{"center", "$maxMatchjob.user.center"},
				{"team", "$maxMatchjob.user.team"},
				{"account_id", "$maxMatchjob.user.account_id"},
			}},
		},
		{
			{"$match", bson.D{
				{"start_date", bson.D{{"$lte", date_nPlus}}},
				{"account_id", bson.D{{"$exists", true}}},
				{"account_id", bson.D{{"$ne", nil}}},
				{"$or", bson.A{
					bson.D{{"last_active_date", nil}},
					bson.D{{"last_active_date", bson.D{{"$gt", date_nPlus}}}},
				}},
				{"team", bson.D{{"$in", bson.A{"Dev", "Tester", "UXUI", "Data Sci", "IT Infra", "DevOps"}}}},
			}},
		},
		{
			{"$group", bson.D{
				{"_id", "$center"},
				{"mapcenter", bson.D{{"$push", "$$ROOT"}}},
			}},
		},
		{
			{"$match", bson.D{
				{"_id", center},
			}},
		},
		{
			{"$unwind", "$mapcenter"},
		},
		{
			{
				"$match", bson.D{
					{"mapcenter.available", status},
				},
			},
		},
		{
			{"$unset", "_id"},
		},
		{{"$replaceRoot", bson.D{{"newRoot", "$mapcenter"}}}},
	}

	query_result, err := collection.Aggregate(r.mgConn.Ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var result []models.StaffCenterStatus

	if err := query_result.All(r.mgConn.Ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}
