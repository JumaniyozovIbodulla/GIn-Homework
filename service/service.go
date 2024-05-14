package service

import (
	"backend_course/lms/pkg/logger"
	"backend_course/lms/storage"
)

type IServiceManager interface {
	Student() studentService
	Teacher() teacherService
	Subjects() subjectsService
	Time() timeService
	Auth() authService
}

type Service struct {
	studentService  studentService
	teacherService  teacherService
	subjectsService subjectsService
	timeService     timeService
	authService     authService
	logger          logger.ILogger
}

func New(storage storage.IStorage, logger logger.ILogger) Service {
	services := Service{}
	services.studentService = NewStudentService(storage, logger)
	services.teacherService = NewTeacherService(storage, logger)
	services.subjectsService = NewSubjectService(storage, logger)
	services.timeService = NewTimeService(storage, logger)
	services.authService = NewAuthService(storage, logger)
	services.logger = logger

	return services
}

func (s Service) Student() studentService {
	return s.studentService
}

func (s Service) Teacher() teacherService {
	return s.teacherService
}

func (s Service) Subjects() subjectsService {
	return s.subjectsService
}

func (s Service) Time() timeService {
	return s.timeService
}

func (s Service) Auth() authService {
	return s.authService
}
