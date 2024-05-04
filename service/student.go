package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/storage"
	"fmt"
)

type studentService struct {
	storage storage.IStorage
}

func NewStudentService(storage storage.IStorage) studentService {
	return studentService{storage: storage}
}

func (s studentService) Create(student models.Student) (string, error) {
	// business logic
	id, err := s.storage.StudentStorage().Create(student)
	if err != nil {
		fmt.Println("error while creating student, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}
