package api

import (
	"backend_course/lms/api/handler"
	"backend_course/lms/storage"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)
// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
func New(store storage.IStorage) *gin.Engine {
	h := handler.NewStrg(store)

	r := gin.Default()

	r.POST("/student", h.CreateStudent)
	r.PUT("/student/:id", h.UpdateStudent)
	r.PATCH("/student/:id", h.UpdateStudentStatus)
	r.GET("/students", h.GetAllStudents)
	r.DELETE("/student/:id", h.DeleteStudent)
	r.GET("/student/:id", h.GetStudent)
	r.POST("/teacher", h.CreateTeacher)
	r.PUT("/teacher/:id", h.UpdateTeacher)
	r.GET("/teachers", h.GetAllTeachers)
	r.DELETE("/teacher/:id", h.DeleteTeacher)
	r.GET("/teacher/:id", h.GetTeacher)
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
