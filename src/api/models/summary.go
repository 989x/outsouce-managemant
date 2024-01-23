package models

import "time"

type SummaryResponse struct {
	ID             string      `json:"_id,omitempty" bson:"_id,omitempty"`
	FinishJobsDate interface{} `json:"finish_jobs_date"`
	AcceptJob      bool        `json:"accept_job"`
	UserID         string      `json:"user_id,omitempty" bson:"user_id,omitempty"`
	JobID          interface{} `json:"job_id"`
	StartJobsDate  time.Time   `json:"start_jobs_date" bson:"start_jobs_date,omitempty"`
	Status         string      `json:"status" bson:"status,omitempty"`
	Available      string      `json:"available" bson:"available,omitempty"`
	Outsource      string      `json:"outsource" bson:"outsource,omitempty"`
	Matchjob       string      `json:"matchjob" bson:"matchjob,omitempty"`
	AddressOnsite  string      `json:"address_onsite" bson:"address_onsite,omitempty"`
	StatusSite     string      `json:"status_site" bson:"status_site,omitempty"`
	Note           string      `json:"note" bson:"note,omitempty"`
	CreatedAt      time.Time   `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt      time.Time   `json:"updatedAt" bson:"createdAt,omitempty"`
	V              int         `json:"__v"`
}
