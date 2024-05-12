package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/storage"
	"context"
	"log"
)

type subjectsService struct {
	storage storage.IStorage
}

func NewSubjectService(storage storage.IStorage) subjectsService {
	return subjectsService{storage: storage}
}

func (s subjectsService) Create(ctx context.Context, subject models.AddSubject) (string, error) {
	// business logic
	id, err := s.storage.SubjectsStorage().Create(ctx, subject)
	if err != nil {
		log.Fatal("error while creating a subject, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}

func (s subjectsService) Update(ctx context.Context, subject models.Subjects) (string, error) {
	// business logic
	id, err := s.storage.SubjectsStorage().Update(ctx, subject)
	if err != nil {
		log.Fatal("error while updating a subject, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}


func (s subjectsService) Delete(ctx context.Context, id string) error {
	err := s.storage.SubjectsStorage().Delete(ctx, id)

	if err != nil {
		log.Fatal("error while deleting a subject: ", err)
		return err
	}

	return nil
}

func (s subjectsService) GetAll(ctx context.Context, req models.GetAllSubjectsRequest) (models.GetAllSubjectsResponse, error) {
	res, err := s.storage.SubjectsStorage().GetAll(ctx, req)
	if err != nil {
		log.Fatal("error while getting all subjects: ", err)
		return res, err
	}

	return res, nil
}

func (s subjectsService) GetSubject(ctx context.Context, id string) (models.Subjects, error) {
	subject, err := s.storage.SubjectsStorage().GetSubject(ctx, id)

	if err != nil {
		log.Fatal("error getting a subject: ", err)
		return subject, err
	}

	return subject, nil
}
