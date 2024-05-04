package models

import "time"

type Teacher struct {
	Id           string    `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	SubjectId    string    `json:"subject_id"`
	StartWorking time.Time `json:"start_working"`
	Phone        string    `json:"phone"`
	Email        string    `json:"mail"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}



type GetAllTeachersRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}


type GetAllTeachersResponse struct {
	Teachers []Teacher `json:"students"`
	Count    int64        `json:"count"`
}