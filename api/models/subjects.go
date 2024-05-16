package models

type Subjects struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AddSubject struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
}

type UpdateSubjects struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
}


type GetAllSubjectsRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

type GetAllSubjectsResponse struct {
	Subjects []Subjects `json:"subjects"`
	Count    int64        `json:"count"`
}
