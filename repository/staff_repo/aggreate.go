package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DashboardPipeline() mongo.Pipeline {
	date_nPlus := time.Now().Add(-24 * time.Hour)

	searchData := bson.M{}
	leaveFilter := bson.M{}
	removeTeam := bson.M{}

	pipeline := mongo.Pipeline{
		{{"$match", bson.D{
			{"start_jobs_date", bson.D{
				{"$lte", date_nPlus},
			}},
			{"$or", bson.A{
				bson.D{{"finish_jobs_date", nil}},
				bson.D{{"finish_jobs_date", bson.D{{"$gte", date_nPlus}}}},
			}},
		}}},
		{{"$group", bson.D{
			{"_id", "$user_id"},
			{"matchjob", bson.D{{"$push", "$$ROOT"}}},
		}}},
		{{"$unwind", "$matchjob"}},
		{{"$sort", bson.D{
			{"matchjob.available", -1},
			{"matchjob.status", -1},
			{"matchjob.start_jobs_date", -1},
			{"matchjob.updatedAt", -1},
			{"matchjob._id", -1},
		}}},
		{{"$group", bson.D{
			{"_id", "$_id"},
			{"maxMatchjob", bson.D{{"$first", "$matchjob"}}},
		}}},
		{{"$unset", "_id"}},
		{{"$lookup", bson.D{
			{"from", "staffs"},
			{"localField", "maxMatchjob.user_id"},
			{"foreignField", "_id"},
			{"as", "maxMatchjob.user"},
		}}},
		{{"$unwind", "$maxMatchjob.user"}},
		{{"$project", bson.D{
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
			{"job_id", "$maxMatchjob.job_id"},
			{"createdAt", "$maxMatchjob.createdAt"},
			{"updatedAt", "$maxMatchjob.updatedAt"},
			// from user
			{"id", "$maxMatchjob.user.id"},
			{"fname", "$maxMatchjob.user.fname"},
			{"lname", "$maxMatchjob.user.lname"},
			{"nname", "$maxMatchjob.user.nname"},
			{"start_date", "$maxMatchjob.user.start_date"},
			{"skill", "$maxMatchjob.user.skill"},
			{"email", "$maxMatchjob.user.email"},
			{"phone", "$maxMatchjob.user.phone"},
			{"resign_description", "$maxMatchjob.user.resign_description"},
			{"active", "$maxMatchjob.user.active"},
			{"isTransfer", "$maxMatchjob.user.isTransfer"},
			{"last_active_date", "$maxMatchjob.user.last_active_date"},
			{"center", "$maxMatchjob.user.center"},
			{"team", "$maxMatchjob.user.team"},
			{"account_id", "$maxMatchjob.user.account_id"},
		}}},
		{{"$match", bson.D{
			{"start_date", bson.D{
				{"$lte", date_nPlus},
			}},
			{"account_id", bson.D{
				{"$exists", true},
				{"$ne", nil},
			}},
			{"$or", bson.A{
				bson.D{{"last_active_date", nil}},
				bson.D{{"last_active_date", bson.D{{"$gt", date_nPlus}}}},
			}},
		}}},
		{{"$match", searchData}},
		{{"$match", leaveFilter}},
		{{"$match", removeTeam}},
	}

	return pipeline
}
