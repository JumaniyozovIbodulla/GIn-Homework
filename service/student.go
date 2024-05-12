package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/storage"
	"context"
	"log"
)

type studentService struct {
	storage storage.IStorage
}

func NewStudentService(storage storage.IStorage) studentService {
	return studentService{storage: storage}
}

func (s studentService) Create(ctx context.Context, student models.AddStudent) (string, error) {
	// business logic
	id, err := s.storage.StudentStorage().Create(ctx, student)
	if err != nil {
		log.Fatal("error while creating a student, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}

func (s studentService) Update(ctx context.Context, student models.Student) (string, error) {
	// business logic
	id, err := s.storage.StudentStorage().Update(ctx, student)
	if err != nil {
		log.Fatal("error while updating a student, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}

func (s studentService) UpdateStatus(ctx context.Context, student models.Student) (string, error) {
	// business logic
	id, err := s.storage.StudentStorage().UpdateStatus(ctx, student)
	if err != nil {
		log.Fatal("error while updating student's status, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}

func (s studentService) Delete(ctx context.Context, id string) error {
	err := s.storage.StudentStorage().Delete(ctx, id)

	if err != nil {
		log.Fatal("error while deleting a student: ", err)
		return err
	}

	return nil
}

func (s studentService) GetAll(ctx context.Context, req models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error) {
	res, err := s.storage.StudentStorage().GetAll(ctx, req)
	if err != nil {
		log.Fatal("error while getting all students: ", err)
		return res, err
	}

	return res, nil
}

func (s studentService) GetStudent(ctx context.Context, id string) (models.GetStudent, error) {
	student, err := s.storage.StudentStorage().GetStudent(ctx, id)

	if err != nil {
		log.Fatal("error getting a student: ", err)
		return student, err
	}

	return student, nil
}