package models

type Teacher struct {
	Id           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	SubjectId    string `json:"subject_id"`
	StartWorking string `json:"start_working"`
	Phone        string `json:"phone"`
	Email        string `json:"mail"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type AddTeacher struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	SubjectId    string `json:"subject_id"`
	StartWorking string `json:"start_working"`
	Phone        string `json:"phone"`
	Email        string `json:"mail"`
	Password     string `json:"password,omitempty"`
}

type GetAllTeachersRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

type GetAllTeachersResponse struct {
	Teachers []Teacher `json:"teachers"`
	Count    int64     `json:"count"`
}
