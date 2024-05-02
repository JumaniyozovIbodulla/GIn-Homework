package api

import (
	"backend_course/lms/api/handler"
	"backend_course/lms/storage"

	"github.com/gin-gonic/gin"
)

func New(store storage.IStorage) *gin.Engine {
	h := handler.NewStrg(store)

	r := gin.Default()

	r.POST("/student", h.CreateStudent)
	r.PUT("/student/:id", h.UpdateStudent)
	r.GET("/students", h.GetAllStudents)
	r.DELETE("/student:id", h.DeleteStudent)
	r.GET("/student:id", h.GetStudent)

	return r

}
