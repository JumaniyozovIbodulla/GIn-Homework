package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/storage"
	"context"
	"log"
)

type timeService struct {
	storage storage.IStorage
}

func NewTimeService(storage storage.IStorage) timeService {
	return timeService{storage: storage}
}

func (s timeService) Create(ctx context.Context, time models.Time) (string, error) {
	// business logic
	id, err := s.storage.TimeStorage().Create(ctx, time)
	if err != nil {
		log.Fatal("error while creating a time table, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}

func (s timeService) Update(ctx context.Context, time models.Time) (string, error) {
	// business logic
	id, err := s.storage.TimeStorage().Update(ctx, time)
	if err != nil {
		log.Fatal("error while updating a time table, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}


func (s timeService) Delete(ctx context.Context, id string) error {
	err := s.storage.TimeStorage().Delete(ctx, id)

	if err != nil {
		log.Fatal("error while deleting a time table: ", err)
		return err
	}

	return nil
}

func (s timeService) GetAll(ctx context.Context, req models.GetAllTimeRequest) (models.GetAllTimeResponse, error) {
	res, err := s.storage.TimeStorage().GetAll(ctx, req)
	if err != nil {
		log.Fatal("error while getting all time tables: ", err)
		return res, err
	}

	return res, nil
}

func (s timeService) GetTimeTable(ctx context.Context, id string) (models.Time, error) {
	time, err := s.storage.TimeStorage().GetTime(ctx, id)

	if err != nil {
		log.Fatal("error getting a time: ", err)
		return time, err
	}
	return time, nil
}
