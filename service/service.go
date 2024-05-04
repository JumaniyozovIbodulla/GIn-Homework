package service

import "backend_course/lms/storage"

type IServiceManager interface {
	Student() studentService
}

type Service struct {
	studentService studentService
}

func New(storage storage.IStorage) Service {
	services := Service{}
	services.studentService = NewStudentService(storage)

	return services
}

func (s Service) Student() studentService {
	return s.studentService
}
