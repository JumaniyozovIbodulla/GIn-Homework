package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg/logger"
	"backend_course/lms/storage"
	"context"
)

type timeService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewTimeService(storage storage.IStorage, logger logger.ILogger) timeService {
	return timeService{
		storage: storage,
		logger: logger,
	}
}

func (s timeService) Create(ctx context.Context, time models.Time) (string, error) {
	id, err := s.storage.TimeStorage().Create(ctx, time)
	if err != nil {
		s.logger.Error("failed to create a time table: ", logger.Error(err))
		return "", err
	}
	return id, nil
}

func (s timeService) Update(ctx context.Context, time models.Time) (string, error) {
	id, err := s.storage.TimeStorage().Update(ctx, time)
	if err != nil {
		s.logger.Error("failed to update a time table: ", logger.Error(err))
		return "", err
	}
	return id, nil
}

func (s timeService) Delete(ctx context.Context, id string) error {
	err := s.storage.TimeStorage().Delete(ctx, id)
	if err != nil {
		s.logger.Error("failed to delete a time table: ", logger.Error(err))
		return err
	}

	return nil
}

func (s timeService) GetAll(ctx context.Context, req models.GetAllTimeRequest) (models.GetAllTimeResponse, error) {
	res, err := s.storage.TimeStorage().GetAll(ctx, req)
	if err != nil {
		s.logger.Error("failed to get all time tables: ", logger.Error(err))
		return res, err
	}
	return res, nil
}

func (s timeService) GetTimeTable(ctx context.Context, id string) (models.Time, error) {
	time, err := s.storage.TimeStorage().GetTime(ctx, id)
	if err != nil {
		s.logger.Error("failed to get a time table: ", logger.Error(err))
		return time, err
	}
	return time, nil
}