package helpers

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ___________________________________>>>>>>>>>>>>>>>>>>>>>>>>>STAFF<<<<<<<<<<<<<<<<<<<<<<<<<___________________________________//
func StaffPrimetiveFilter(available []interface{}, team []interface{}, status []interface{}, center []interface{}) primitive.D {
	filter := bson.D{
		{"$and", bson.A{
			bson.D{{"available", bson.M{"$in": available}}},
			bson.D{{"team", bson.M{"$in": team}}},
			bson.D{{"matchjob", bson.M{"$nin": status}}},
		}},
	}
	return filter
}

func StaffPipeLineTotal(fillter primitive.D, date_nPlus time.Time) []primitive.D {
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
	return pipeline
}

func GetCountCenterStaff(center string, status string, date_nPlus time.Time) []primitive.D {
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
	return pipeline
}

func GetProjectTotal(date_nPlus time.Time) []primitive.D {
	pipeline := []bson.D{
		{{"$match", bson.D{
			{"start_jobs_date", bson.D{{"$lte", date_nPlus}}},
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
			{"matchjob", "$maxMatchjob.matchjob"},
			{"address_onsite", "$maxMatchjob.address_onsite"},
			{"status_site", "$maxMatchjob.status_site"},
			{"createdAt", "$maxMatchjob.createdAt"},
			{"updatedAt", "$maxMatchjob.updatedAt"},
			{"available", "$maxMatchjob.available"},
			{"outsource", "$maxMatchjob.outsource"},
			{"note", "$maxMatchjob.note"},
			{"job_id", "$maxMatchjob.job_id"},
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
		}}},
		{{"$match", bson.D{
			{"start_date", bson.D{{"$lte", date_nPlus}}},
			{"account_id", bson.D{
				{"$exists", true},
				{"$ne", nil},
			}},
			{"$or", bson.A{
				bson.D{{"last_active_date", nil}},
				bson.D{{"last_active_date", bson.D{{"$gt", date_nPlus}}}},
			}},
			{"team", bson.D{{"$in", bson.A{
				"Dev", "Tester", "UXUI", "Data Sci", "IT Infra", "DevOps",
			}}}},
		}}},
		{{"$group", bson.D{
			{"_id", "$job_id"},
			{"getproject", bson.D{{"$push", "$$ROOT"}}},
		}}},
		{{"$addFields", bson.D{
			{"arraySize", bson.D{{"$size", "$getproject"}}},
		}}},
		{{"$addFields", bson.D{
			{"projectName", bson.D{{"$arrayElemAt", bson.A{"$getproject.matchjob", 0}}}},
		}}},
		{{"$project", bson.D{
			{"projectName", "$projectName"},
			{"projectParticipant", "$arraySize"},
		}}},
		{{"$sort", bson.D{
			{"projectParticipant", -1},
		}}},
	}
	return pipeline
}

func GetStaffParticipant(date_nPlus time.Time, jobId primitive.ObjectID) []primitive.D {
	pipeline := []bson.D{
		{{"$match", bson.D{
			{"start_jobs_date", bson.D{{"$lte", date_nPlus}}},
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
			{"matchjob", "$maxMatchjob.matchjob"},
			{"address_onsite", "$maxMatchjob.address_onsite"},
			{"status_site", "$maxMatchjob.status_site"},
			{"createdAt", "$maxMatchjob.createdAt"},
			{"updatedAt", "$maxMatchjob.updatedAt"},
			{"available", "$maxMatchjob.available"},
			{"outsource", "$maxMatchjob.outsource"},
			{"note", "$maxMatchjob.note"},
			{"job_id", "$maxMatchjob.job_id"},
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
		}}},
		{{"$match", bson.D{
			{"start_date", bson.D{{"$lte", date_nPlus}}},
			{"account_id", bson.D{{"$exists", true}}},
			{"account_id", bson.D{{"$ne", nil}}},
			{"$or", bson.A{
				bson.D{{"last_active_date", nil}},
				bson.D{{"last_active_date", bson.D{{"$gt", date_nPlus}}}},
			}},
			{"team", bson.D{{"$in", bson.A{"Dev", "Tester", "UXUI", "Data Sci", "IT Infra", "DevOps"}}}},
		}}},
		{{"$group", bson.D{
			{"_id", "$job_id"},
			{"getproject", bson.D{{"$push", "$$ROOT"}}},
		}}},
		{{"$match", bson.D{
			{"_id", jobId},
		}}},
		{{"$unset", "_id"}},
		{{"$unwind", "$getproject"}},
		{{"$replaceRoot", bson.D{
			{"newRoot", "$getproject"},
		}}},
	}
	return pipeline
}

func GetAllStaff(date_nPlus time.Time) []primitive.D {
	pipeline := []bson.D{
		{{"$match", bson.D{
			{"start_jobs_date", bson.D{{"$lte", date_nPlus}}},
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
			{"matchjob", "$maxMatchjob.matchjob"},
			{"address_onsite", "$maxMatchjob.address_onsite"},
			{"status_site", "$maxMatchjob.status_site"},
			{"createdAt", "$maxMatchjob.createdAt"},
			{"updatedAt", "$maxMatchjob.updatedAt"},
			{"available", "$maxMatchjob.available"},
			{"outsource", "$maxMatchjob.outsource"},
			{"note", "$maxMatchjob.note"},
			{"job_id", "$maxMatchjob.job_id"},
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
			{"skill", "$maxMatchjob.user.skill"},
		}}},
		{{"$match", bson.D{
			{"start_date", bson.D{{"$lte", date_nPlus}}},
			{"account_id", bson.D{{"$exists", true}}},
			{"account_id", bson.D{{"$ne", nil}}},
			{"$or", bson.A{
				bson.D{{"last_active_date", nil}},
				bson.D{{"last_active_date", bson.D{{"$gt", date_nPlus}}}},
			}},
			{"team", bson.D{{"$in", bson.A{"Dev", "Tester", "UXUI", "Data Sci", "IT Infra", "DevOps"}}}},
		}}},
	}
	return pipeline
}

func GetSearchStaff(fillter primitive.D, date_nPlus time.Time) []primitive.D {
	pipeline := []bson.D{
		{{"$match", bson.D{
			{"start_jobs_date", bson.D{{"$lte", date_nPlus}}},
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
			{"matchjob", "$maxMatchjob.matchjob"},
			{"address_onsite", "$maxMatchjob.address_onsite"},
			{"status_site", "$maxMatchjob.status_site"},
			{"createdAt", "$maxMatchjob.createdAt"},
			{"updatedAt", "$maxMatchjob.updatedAt"},
			{"available", "$maxMatchjob.available"},
			{"outsource", "$maxMatchjob.outsource"},
			{"note", "$maxMatchjob.note"},
			{"job_id", "$maxMatchjob.job_id"},
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
			{"skill", "$maxMatchjob.user.skill"},
		}}},
		{{"$match", bson.D{
			{"start_date", bson.D{{"$lte", date_nPlus}}},
			{"account_id", bson.D{{"$exists", true}}},
			{"account_id", bson.D{{"$ne", nil}}},
			{"$or", bson.A{
				bson.D{{"last_active_date", nil}},
				bson.D{{"last_active_date", bson.D{{"$gt", date_nPlus}}}},
			}},
			{"team", bson.D{{"$in", bson.A{"Dev", "Tester", "UXUI", "Data Sci", "IT Infra", "DevOps"}}}},
		}}},
		{
			{"$match", fillter},
		},
	}
	return pipeline
}

// ___________________________________>>>>>>>>>>>>>>>>>>>>>>>>>STAFF<<<<<<<<<<<<<<<<<<<<<<<<<___________________________________//
