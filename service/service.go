package service

import "backend_course/lms/storage"

type IServiceManager interface {
	Student() studentService
	Teacher() teacherService
	Subjects() subjectsService
	Time() timeService
}

type Service struct {
	studentService studentService
	teacherService teacherService
	subjectsService subjectsService
	timeService timeService
}

func New(storage storage.IStorage) Service {
	services := Service{}
	services.studentService = NewStudentService(storage)
	services.teacherService = NewTeacherService(storage)
	services.subjectsService = NewSubjectService(storage)
	services.timeService = NewTimeService(storage)

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