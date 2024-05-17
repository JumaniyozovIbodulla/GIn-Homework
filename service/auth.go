package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/config"
	"backend_course/lms/pkg"
	"backend_course/lms/pkg/check"
	"backend_course/lms/pkg/jwt"
	"backend_course/lms/pkg/logger"
	"backend_course/lms/storage"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cast"
)

type authService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewAuthService(storage storage.IStorage, logger logger.ILogger) authService {
	return authService{
		storage: storage,
		logger:  logger,
	}
}

func (s authService) Login(ctx context.Context, req models.LoginRequest) (models.LoginResponse, error) {
	resp := models.LoginResponse{}

	teacher, err := s.storage.TeacherStorage().GetTeacherByLogin(ctx, req.Login)
	if err != nil {
		s.logger.Error("failed to get teacher by login: ", logger.Error(err))
		return resp, err
	}

	if err = pkg.CompareHashAndPassword(teacher.Password, req.Password); err != nil {
		s.logger.Error("password is not match: ", logger.Error(err))
		return resp, errors.New("password doesn't match")
	}

	m := make(map[interface{}]interface{})
	m["user_id"] = teacher.Id
	m["user_role"] = config.TEACHER_TYPE
	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		s.logger.Error("failed to get access and refresh token: ", logger.Error(err))
		return resp, err
	}
	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken

	return resp, nil
}

func (s authService) TeacherRegister(ctx context.Context, req models.RegisterRequest) error {
	exists := s.storage.TeacherStorage().IsTeacherExists(ctx, req.Mail)
	if exists {
		otp := pkg.GenerateOTP()
		msg := fmt.Sprintf("Your code: %v. DON'T give anyone", otp)
		err := s.storage.Redis().SetX(ctx, req.Mail, otp, time.Minute*2)
		if err != nil {
			return err
		}

		err = check.SendEmail(req.Mail, msg)
		if err != nil {
			return err
		}
	} else {
		s.logger.Error("failed to get teacher by login: ", logger.Error(errors.New("something went wrong")))
		return errors.New("email already exists")
	}
	return nil
}

func (s authService) TeacherRegisterConfirm(ctx context.Context, req models.RegisterConfirmRequest) error {
	code := s.storage.Redis().Get(ctx, req.AddTeacher.Email)
	resultCode := cast.ToInt(code)
	if resultCode == req.Code {
		_, err := s.storage.TeacherStorage().Create(ctx, req.AddTeacher)

		if err != nil {
			s.logger.Error("failed to create a new teacher: ", logger.Error(err))
			return err
		}
		return nil
	}
	s.logger.Error("code is not match or expired code: ")
	return errors.New("code is not match or expired code")
}


func (s authService) TeacherOTP(ctx context.Context, req models.RegisterConfirmRequest) error {
	code := s.storage.Redis().Get(ctx, req.AddTeacher.Email)
	resultCode := cast.ToInt(code)
	if resultCode == req.Code {
		_, err := s.storage.TeacherStorage().Create(ctx, req.AddTeacher)

		if err != nil {
			s.logger.Error("failed to create a new teacher: ", logger.Error(err))
			return err
		}
		return nil
	}
	s.logger.Error("code is not match or expired code: ")
	return errors.New("code is not match or expired code")
}