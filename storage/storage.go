package storage

import "backend_course/lms/api/models"

type IStorage interface {
	CloseDB()
	StudentStorage() StudentStorage
	TeacherStorage() TeacherStorage
}

type StudentStorage interface {
	Create(student models.Student) (string, error)
	Update(student models.Student) (string, error)
	UpdateStatus(student models.Student) (string, error)
	Delete(id string) (error)
	GetStudent(id string) (models.GetStudent, error)
	GetAll(req models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error)
}

type TeacherStorage interface {
	Create(teacher models.Teacher) (string, error)
	Update(teacher models.Teacher) (string, error)
	Delete(id string) (error)
	GetTeacher(id string) (models.Teacher, error)
	GetAll(req models.GetAllTeachersRequest) (models.GetAllTeachersResponse, error)
}