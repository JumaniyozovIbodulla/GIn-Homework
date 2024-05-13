package models

type Student struct {
	Id         string `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
	ExternalId string `json:"external_id"`
	Phone      string `json:"phone"`
	Email      string `json:"mail"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	IsActive   bool   `json:"is_active"`
}

type CheckLessonStudent struct {
	StudentId   string  `json:"student_id"`
	StudentName string  `json:"student_name"`
	StudentAge  uint16  `json:"student_age"`
	SubjectName string  `json:"subject_name"`
	TeacherName string  `json:"teacher_name"`
	RoomName    string  `json:"room_name"`
	TimeLeft    float64 `json:"time_left"`
}

type AddStudent struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
	ExternalId string `json:"external_id"`
	Phone      string `json:"phone"`
	Email      string `json:"mail"`
	IsActive   bool   `json:"is_active"`
	Password   string `json:"password,omitempty"`
}

type GetStudent struct {
	Id         string `json:"id"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Age        int    `json:"age,omitempty"`
	ExternalId string `json:"external_id,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Email      string `json:"email,omitempty"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	IsActive   bool   `json:"is_active"`
}

type GetAllStudentsRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

type GetAllStudentsResponse struct {
	Students []GetStudent `json:"students"`
	Count    int64        `json:"count"`
}
