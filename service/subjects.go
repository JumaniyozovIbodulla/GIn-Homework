package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg/logger"
	"backend_course/lms/storage"
	"context"
)

type subjectsService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewSubjectService(storage storage.IStorage, logger logger.ILogger) subjectsService {
	return subjectsService{
		storage: storage,
		logger: logger,
	}
}

func (s subjectsService) Create(ctx context.Context, subject models.AddSubject) (string, error) {
	id, err := s.storage.SubjectsStorage().Create(ctx, subject)
	if err != nil {
		s.logger.Error("failed to create a subject: ", logger.Error(err))
		return "", err
	}
	return id, nil
}

func (s subjectsService) Update(ctx context.Context, subject models.Subjects) (string, error) {
	id, err := s.storage.SubjectsStorage().Update(ctx, subject)
	if err != nil {
		s.logger.Error("failed to update a subject: ", logger.Error(err))
		return "", err
	}
	return id, nil
}

func (s subjectsService) Delete(ctx context.Context, id string) error {
	err := s.storage.SubjectsStorage().Delete(ctx, id)
	if err != nil {
		s.logger.Error("failed to delete a subject: ", logger.Error(err))
		return err
	}

	return nil
}

func (s subjectsService) GetAll(ctx context.Context, req models.GetAllSubjectsRequest) (models.GetAllSubjectsResponse, error) {
	res, err := s.storage.SubjectsStorage().GetAll(ctx, req)
	if err != nil {
		s.logger.Error("failed to get all subjects: ", logger.Error(err))
		return res, err
	}

	return res, nil
}

func (s subjectsService) GetSubject(ctx context.Context, id string) (models.Subjects, error) {
	subject, err := s.storage.SubjectsStorage().GetSubject(ctx, id)
	if err != nil {
		s.logger.Error("failed to get a subject: ", logger.Error(err))
		return subject, err
	}

	return subject, nil
}
