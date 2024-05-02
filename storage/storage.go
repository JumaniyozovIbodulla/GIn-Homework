package storage

import "backend_course/lms/api/models"

type IStorage interface {
	CloseDB()
	StudentStorage() StudentStorage
}

type StudentStorage interface {
	Create(student models.Student) (string, error)
	Update(student models.Student) (string, error)
	Delete(student models.Student) (string, error)
	GetStudent(student models.GetStudent) (models.GetStudent, error)
	GetAll(req models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error)
}
