package api

import (
	"backend_course/lms/api/handler"
	"backend_course/lms/service"
	"backend_course/lms/storage"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
func New(store storage.IStorage, service service.IServiceManager) *gin.Engine {
	h := handler.NewStrg(store, service)

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
	r.POST("/subject", h.CreateSubject)
	r.PUT("/subject/:id", h.UpdateSubject)
	r.DELETE("/subject/:id", h.DeleteSubject)
	r.GET("/subject/:id", h.GetSubject)
	r.GET("/subjects", h.GetAllSubjects)
	r.POST("/time", h.CreateTime)
	r.PUT("/time/:id", h.UpdateTime)
	r.DELETE("/time/:id", h.DeleteTime)
	r.GET("/time/:id", h.GetTime)
	r.GET("/time-tables", h.GetAllTimeTables)
	
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
