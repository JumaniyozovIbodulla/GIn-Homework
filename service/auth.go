package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/config"
	"backend_course/lms/pkg"
	"backend_course/lms/pkg/jwt"
	"backend_course/lms/pkg/logger"
	"backend_course/lms/storage"
	"context"
	"errors"
)

type authService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewAuthService(storage storage.IStorage, logger logger.ILogger) authService {
	return authService{
		storage: storage,
		logger: logger,
	}
}

func (s authService) TeacherLogin(ctx context.Context, req models.LoginRequest) (models.LoginResponse, error) {
	resp := models.LoginResponse{}

	teacher, err := s.storage.TeacherStorage().GetTeacherByLogin(ctx, req.Login)
	if err != nil {
		s.logger.Error("failed to get teacher by login: ", logger.Error(err))
		return resp, err
	}

	if err = pkg.CompareHashAndPassword(teacher.Password, req.Password); err != nil {
		return resp, errors.New("password doesn't match")
	}

	m := make(map[interface{}]interface{})
	m["user_id"] = teacher.Id
	m["user_role"] = config.TEACHER_TYPE
	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		return resp, err
	}
	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken

	return resp, nil
}
