package models

import "time"

type CountStaffDashBoard struct {
	Total   interface{}
	Site    interface{}
	Project interface{}
}

type StaffTotal struct {
	All          int
	AllAvailable int
	AllOnBoard   int
	DevOnBoard   int
	DevAvailable int
	AllDev       int
	ItOnBoard    int
	ItAvailable  int
	AllIt        int
}

type StaffCountCenter struct {
	BnkAvaSlide  interface{}
	BnkOnbSlide  interface{}
	BnkAvaSCount int
	BnkOnbSCount int
	ChmAvaSlide  interface{}
	ChmOnbSlide  interface{}
	ChmAvaSCount int
	ChmOnbSCount int
	KhnAvaSlide  interface{}
	KhnOnbSlide  interface{}
	KhnAvaSCount int
	KhnOnbSCount int
	HdyAvaSlide  interface{}
	HdyOnbSlide  interface{}
	HdyAvaSCount int
	HdyOnbSCount int
}

type StaffGetProject struct {
	ID                 string `json:"_id" bson:"_id,omitempty"`
	ProjectName        string `json:"x" bson:"projectName,omitempty"`
	ProjectParticipant int    `json:"y" bson:"projectParticipant,omitempty"`
}

type StaffProjectApex struct {
	X string `json:"x"`
	Y int    `json:"y"`
}

type StaffDashBoard struct {
	Obj_ID         string      `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID         string      `json:"user_id" bson:"user_id,omitempty"`
	StartJobsDate  time.Time   `json:"start_jobs_date" bson:"start_jobs_date,omitempty"`
	FinishJobsDate time.Time   `json:"finish_jobs_date" bson:"finish_jobs_date,omitempty"`
	Status         string      `json:"status" bson:"status,omitempty"`
	Available      string      `json:"available" bson:"available,omitempty"`
	Outsource      string      `json:"outsource" bson:"outsource,omitempty"`
	Matchjob       string      `json:"matchjob" bson:"matchjob,omitempty"`
	AddressOnsite  string      `json:"address_onsite" bson:"address_onsite,omitempty"`
	StatusSite     string      `json:"status_site" bson:"status_site,omitempty"`
	Note           string      `json:"note" bson:"note,omitempty"`
	CreatedAt      time.Time   `json:"createdAt"  bson:"createdAt,omitempty"`
	UpdatedAt      time.Time   `json:"updatedAt"  bson:"updatedAt,omitempty"`
	ID             string      `json:"id" bson:"id,omitempty"`
	Fname          string      `json:"fname" bson:"fname,omitempty"`
	Lname          string      `json:"lname" bson:"lname,omitempty"`
	Nname          string      `json:"nname" bson:"nname,omitempty"`
	StartDate      time.Time   `json:"start_date" bson:"start_date,omitempty"`
	Active         bool        `json:"active" bson:"active,omitempty"`
	IsTransfer     bool        `json:"isTransfer" bson:"isTransfer,omitempty"`
	LastActiveDate interface{} `json:"last_active_date" bson:"last_active_date,omitempty"`
	Center         string      `json:"center" bson:"center,omitempty"`
	Team           string      `json:"team" bson:"team,omitempty"`
	AccountID      string      `json:"account_id" bson:"account_id,omitempty"`
}

type StaffCenterStatus struct {
	Obj_ID         string      `json:"_id" bson:"_id,omitempty"`
	UserID         string      `json:"user_id" bson:"user_id,omitempty"`
	StartJobsDate  time.Time   `json:"start_jobs_date" bson:"start_jobs_date,omitempty"`
	FinishJobsDate time.Time   `json:"finish_jobs_date" bson:"finish_jobs_date,omitempty"`
	Matchjob       string      `json:"matchjob" bson:"matchjob,omitempty"`
	Status         string      `json:"status" bson:"status,omitempty"`
	AddressOnsite  string      `json:"address_onsite" bson:"address_onsite,omitempty"`
	StatusSite     string      `json:"status_site" bson:"status_site,omitempty"`
	CreatedAt      time.Time   `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt      time.Time   `json:"updatedAt" bson:"updatedAt,omitempty"`
	Available      string      `json:"available" bson:"available,omitempty"`
	Outsource      string      `json:"outsource" bson:"outsource,omitempty"`
	Note           string      `json:"note" bson:"note,omitempty"`
	JobID          string      `json:"job_id" bson:"job_id,omitempty"`
	ID             string      `json:"id" bson:"id,omitempty"`
	Fname          string      `json:"fname" bson:"fname,omitempty"`
	Lname          string      `json:"lname" bson:"lname,omitempty"`
	Nname          string      `json:"nname" bson:"nname,omitempty"`
	StartDate      time.Time   `json:"start_date" bson:"start_date,omitempty"`
	Active         bool        `json:"active" bson:"active,omitempty"`
	IsTransfer     bool        `json:"isTransfer" bson:"isTransfer,omitempty"`
	LastActiveDate interface{} `json:"last_active_date" bson:"last_active_date,omitempty"`
	Center         string      `json:"center" bson:"center,omitempty"`
	Team           string      `json:"team" bson:"team,omitempty"`
	AccountID      string      `json:"account_id" bson:"account_id,omitempty"`
}

type StaffParticipant struct {
	Obj_ID         string      `json:"_id" bson:"_id,omitempty"`
	UserID         string      `json:"user_id" bson:"user_id,omitempty"`
	StartJobsDate  time.Time   `json:"start_jobs_date" bson:"start_jobs_date,omitempty"`
	FinishJobsDate time.Time   `json:"finish_jobs_date" bson:"finish_jobs_date,omitempty"`
	Status         string      `json:"status" bson:"status,omitempty"`
	Matchjob       string      `json:"matchjob" bson:"matchjob,omitempty"`
	AddressOnsite  string      `json:"address_onsite" bson:"address_onsite,omitempty"`
	StatusSite     string      `json:"status_site" bson:"status_site,omitempty"`
	CreatedAt      time.Time   `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt      time.Time   `json:"updatedAt" bson:"updatedAt,omitempty"`
	Available      string      `json:"available" bson:"available,omitempty"`
	Outsource      string      `json:"outsource" bson:"outsource,omitempty"`
	Note           string      `json:"note" bson:"note,omitempty"`
	JobID          string      `json:"job_id" bson:"job_id,omitempty"`
	ID             string      `json:"id" bson:"id,omitempty"`
	Fname          string      `json:"fname" bson:"fname,omitempty"`
	Lname          string      `json:"lname" bson:"lname,omitempty"`
	Nname          string      `json:"nname" bson:"nname,omitempty"`
	StartDate      time.Time   `json:"start_date" bson:"start_date,omitempty"`
	Active         bool        `json:"active" bson:"active,omitempty"`
	IsTransfer     bool        `json:"isTransfer" bson:"isTransfer,omitempty"`
	LastActiveDate interface{} `json:"last_active_date" bson:"last_active_date,omitempty"`
	Center         string      `json:"center" bson:"center,omitempty"`
	Team           string      `json:"team" bson:"team,omitempty"`
	AccountID      string      `json:"account_id" bson:"account_id,omitempty"`
}

type Staff struct {
	Obj_ID         string      `json:"_id" bson:"_id,omitempty"`
	UserID         string      `json:"user_id" bson:"user_id,omitempty"`
	StartJobsDate  time.Time   `json:"start_jobs_date" bson:"start_jobs_date,omitempty"`
	FinishJobsDate time.Time   `json:"finish_jobs_date" bson:"finish_jobs_date,omitempty"`
	Status         string      `json:"status" bson:"status,omitempty"`
	Matchjob       string      `json:"matchjob" bson:"matchjob,omitempty"`
	AddressOnsite  string      `json:"address_onsite" bson:"address_onsite,omitempty"`
	StatusSite     string      `json:"status_site" bson:"status_site,omitempty"`
	CreatedAt      time.Time   `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt      time.Time   `json:"updatedAt" bson:"updatedAt,omitempty"`
	Available      string      `json:"available" bson:"available,omitempty"`
	Outsource      string      `json:"outsource" bson:"outsource,omitempty"`
	Note           string      `json:"note" bson:"note,omitempty"`
	JobID          string      `json:"job_id" bson:"job_id,omitempty"`
	ID             string      `json:"id" bson:"id,omitempty"`
	Fname          string      `json:"fname" bson:"fname,omitempty"`
	Lname          string      `json:"lname" bson:"lname,omitempty"`
	Nname          string      `json:"nname" bson:"nname,omitempty"`
	StartDate      time.Time   `json:"start_date" bson:"start_date,omitempty"`
	Active         bool        `json:"active" bson:"active,omitempty"`
	IsTransfer     bool        `json:"isTransfer" bson:"isTransfer,omitempty"`
	LastActiveDate interface{} `json:"last_active_date" bson:"last_active_date,omitempty"`
	Center         string      `json:"center" bson:"center,omitempty"`
	Team           string      `json:"team" bson:"team,omitempty"`
	AccountID      string      `json:"account_id" bson:"account_id,omitempty"`
}
