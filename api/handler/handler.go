package handler

import (
	"backend_course/lms/api/models"
	"backend_course/lms/config"
	"backend_course/lms/pkg/jwt"
	"backend_course/lms/pkg/logger"
	"backend_course/lms/service"
	"backend_course/lms/storage"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service service.IServiceManager
	Log     logger.ILogger
}

func NewStrg(store storage.IStorage, service service.IServiceManager, log logger.ILogger) Handler {
	return Handler{
		Service: service,
		Log:     log,
	}
}

func handleResponse(c *gin.Context, log logger.ILogger, msg string, statusCode int, data interface{}) {
	resp := models.Response{}

	if statusCode >= 100 && statusCode <= 199 {
		resp.Description = config.ERR_INFORMATION
	} else if statusCode >= 200 && statusCode <= 299 {
		resp.Description = config.SUCCESS
		log.Info("REQUEST SUCCEEDED", logger.Any("msg: ", msg), logger.Int("status: ", statusCode))
	} else if statusCode >= 300 && statusCode <= 399 {
		resp.Description = config.ERR_REDIRECTION
	} else if statusCode >= 400 && statusCode <= 499 {
		resp.Description = config.ERR_BADREQUEST
		log.Error("BAD REQUEST", logger.Any("error: ", msg), logger.Any("data: ", data))
	} else {
		resp.Description = config.ERR_INTERNAL_SERVER
		log.Error("ERR_INTERNAL_SERVER", logger.Any("error: ", msg))
	}

	resp.StatusCode = statusCode
	resp.Data = data

	c.JSON(resp.StatusCode, resp)
}

func ParsePageQueryParam(c *gin.Context) (uint64, error) {
	pageStr := c.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.ParseUint(pageStr, 10, 30)
	if err != nil {
		return 0, err
	}
	if page == 0 {
		return 1, nil
	}
	return page, nil
}

func ParseLimitQueryParam(c *gin.Context) (uint64, error) {
	limitStr := c.Query("limit")
	if limitStr == "" {
		limitStr = "10"
	}
	limit, err := strconv.ParseUint(limitStr, 10, 30)
	if err != nil {
		return 0, err
	}

	if limit == 0 {
		return 10, nil
	}
	return limit, nil
}

func getAuthInfo(c *gin.Context) (models.AuthInfo, error) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		return models.AuthInfo{}, errors.New("unauthorized")
	}

	m, err := jwt.ExtractClaims(accessToken)
	if err != nil {
		return models.AuthInfo{}, err
	}

	fmt.Println("m: ", m)
	role := m["user_role"].(string)
	if !(role == config.TEACHER_TYPE || role == config.STUDENT_TYPE) {
		return models.AuthInfo{}, errors.New("unauthorized")
	}

	return models.AuthInfo{
		UserID:   m["user_id"].(string),
		UserRole: role,
	}, nil
}
