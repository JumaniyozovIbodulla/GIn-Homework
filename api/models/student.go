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
	StudentName string  `json:"student_name"`
	StudentAge  uint16  `json:"student_age"`
	SubjectName string  `json:"subject_name"`
	TeacherName string  `json:"teacher_name"`
	RoomName    string  `json:"room_name"`
	TimeLeft    float64 `json:"time_left_in_minutes"`
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

type GetAllStudentsAttandenceReportRequest struct {
	StudentId string `json:"student_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	TeacherId string `json:"teacher_id"`
	Page      uint64 `json:"page"`
	Limit     uint64 `json:"limit"`
}

type GetAllStudentsAttandenceReportResponse struct {
	Students []StudentAttandenceReport `json:"students"`
	Count    int64                     `json:"count"`
}

type StudentAttandenceReport struct {
	StudentId        string  `json:"student_id"`
	StudentName      string  `json:"student_name"`
	StudentCreatedAt string  `json:"student_created_at"`
	TeacherName      string  `json:"teacher_name"`
	StudyTime        float64 `json:"study_time"`
	AvgStudyTime     float64 `json:"avg_study_time"`
}

type UploadStudentImage struct {
	Id   string `json:"id"`
	Path string `json:"image"`
}