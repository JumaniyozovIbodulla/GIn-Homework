package models

type Time struct {
	Id        string `json:"id"`
	TeacherId string `json:"teacher_id"`
	StudentId string `json:"student_id"`
	SubjectId string `json:"subject_id"`
	FromDate  string `json:"from_date"`
	ToDate    string `json:"to_date"`
	RoomName  string `json:"room_name"`
}

type AddTime struct {
	TeacherId string `json:"teacher_id"`
	StudentId string `json:"student_id"`
	SubjectId string `json:"subject_id"`
	FromDate  string `json:"from_date"`
	ToDate    string `json:"to_date"`
	RoomName  string `json:"room_name"`
}

type GetAllTimeRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

type GetAllTimeResponse struct {
	Time  []Time `json:"time_tables"`
	Count int64  `json:"count"`
}
