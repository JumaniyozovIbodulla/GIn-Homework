package storage

import (
	"backend_course/lms/api/models"
	"context"
)

type IStorage interface {
	CloseDB()
	StudentStorage() StudentStorage
	TeacherStorage() TeacherStorage
	SubjectsStorage() SubjectStorage
	TimeStorage() TimeStorage
}

type StudentStorage interface {
	Create(ctx context.Context, student models.AddStudent) (string, error)
	Update(ctx context.Context, student models.Student) (string, error)
	UpdateStatus(ctx context.Context, student models.Student) (string, error)
	Delete(ctx context.Context, id string) error
	GetStudent(ctx context.Context, id string) (models.GetStudent, error)
	GetAll(ctx context.Context, req models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error)
}

type TeacherStorage interface {
	Create(ctx context.Context, teacher models.AddTeacher) (string, error)
	Update(ctx context.Context, teacher models.Teacher) (string, error)
	Delete(ctx context.Context, id string) error
	GetTeacher(ctx context.Context, id string) (models.Teacher, error)
	GetAll(ctx context.Context, req models.GetAllTeachersRequest) (models.GetAllTeachersResponse, error)
}

type SubjectStorage interface {
	Create(ctx context.Context, subject models.AddSubject) (string, error)
	Update(ctx context.Context, subject models.Subjects) (string, error)
	Delete(ctx context.Context, id string) error
	GetSubject(ctx context.Context, id string) (models.Subjects, error)
	GetAll(ctx context.Context, req models.GetAllSubjectsRequest) (models.GetAllSubjectsResponse, error)
}

type TimeStorage interface {
	Create(ctx context.Context, time models.Time) (string, error)
	Update(ctx context.Context, time models.Time) (string, error)
	Delete(ctx context.Context, id string) error
	GetTime(ctx context.Context, id string) (models.Time, error)
	GetAll(ctx context.Context, req models.GetAllTimeRequest) (models.GetAllTimeResponse, error)
}
