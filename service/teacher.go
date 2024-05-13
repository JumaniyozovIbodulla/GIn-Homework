package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/storage"
	"context"
)

type teacherService struct {
	storage storage.IStorage
}

func NewTeacherService(storage storage.IStorage) teacherService {
	return teacherService{storage: storage}
}

func (s teacherService) Create(ctx context.Context, teacher models.AddTeacher) (string, error) {
	// business logic
	id, err := s.storage.TeacherStorage().Create(ctx, teacher)
	if err != nil {
		return "", err
	}
	// business logic
	return id, nil
}

func (s teacherService) Update(ctx context.Context, teacher models.Teacher) (string, error) {
	// business logic
	id, err := s.storage.TeacherStorage().Update(ctx, teacher)
	if err != nil {
		return "", err
	}
	// business logic
	return id, nil
}

func (s teacherService) Delete(ctx context.Context, id string) error {
	err := s.storage.TeacherStorage().Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

func (s teacherService) GetAll(ctx context.Context, req models.GetAllTeachersRequest) (models.GetAllTeachersResponse, error) {
	res, err := s.storage.TeacherStorage().GetAll(ctx, req)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s teacherService) GetTeacher(ctx context.Context, id string) (models.Teacher, error) {
	teacher, err := s.storage.TeacherStorage().GetTeacher(ctx, id)

	if err != nil {
		return teacher, err
	}

	return teacher, nil
}
