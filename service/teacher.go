package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg/logger"
	"backend_course/lms/storage"
	"context"
)

type teacherService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewTeacherService(storage storage.IStorage, logger logger.ILogger) teacherService {
	return teacherService{
		storage: storage,
		logger:  logger,
	}
}

func (s teacherService) Create(ctx context.Context, teacher models.AddTeacher) (string, error) {
	id, err := s.storage.TeacherStorage().Create(ctx, teacher)
	if err != nil {
		s.logger.Error("failed to create a teacher: ", logger.Error(err))
		return "", err
	}
	return id, nil
}

func (s teacherService) Update(ctx context.Context, teacher models.Teacher) (string, error) {
	id, err := s.storage.TeacherStorage().Update(ctx, teacher)
	if err != nil {
		s.logger.Error("failed to create a teacher: ", logger.Error(err))
		return "", err
	}
	return id, nil
}

func (s teacherService) Delete(ctx context.Context, id string) error {
	err := s.storage.TeacherStorage().Delete(ctx, id)
	if err != nil {
		s.logger.Error("failed to delete a teacher: ", logger.Error(err))
		return err
	}

	return nil
}

func (s teacherService) GetAll(ctx context.Context, req models.GetAllTeachersRequest) (models.GetAllTeachersResponse, error) {
	res, err := s.storage.TeacherStorage().GetAll(ctx, req)
	if err != nil {
		s.logger.Error("failed to delete all teachers: ", logger.Error(err))
		return res, err
	}
	return res, nil
}

func (s teacherService) GetTeacher(ctx context.Context, id string) (models.Teacher, error) {
	teacher, err := s.storage.TeacherStorage().GetTeacher(ctx, id)
	if err != nil {
		s.logger.Error("failed to get a teacher: ", logger.Error(err))
		return teacher, err
	}

	return teacher, nil
}

func (s teacherService) CheckTeacherLesson(ctx context.Context, id string) (models.CheckLessonTeacher, error) {
	checkTeacher, err := s.storage.TeacherStorage().CheckTeacherLesson(ctx, id)
	if err != nil {
		s.logger.Error("failed to get a teacher's lesson: ", logger.Error(err))
		return checkTeacher, err
	}
	return checkTeacher, nil
}