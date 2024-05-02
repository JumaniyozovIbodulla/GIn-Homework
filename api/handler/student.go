package handler

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg/check"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


func (h Handler) CreateStudent(c *gin.Context) {
	student := models.Student{}

	if err := c.ShouldBindJSON(&student); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	if err := check.ValidateYear(student.Age); err != nil {
		handleResponse(c, "error while validating student age, year: "+strconv.Itoa(student.Age), http.StatusBadRequest, err.Error())

		return
	}

	id, err := h.Store.StudentStorage().Create(student)
	if err != nil {
		handleResponse(c, "error while creating student", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Created successfully", http.StatusOK, id)
}

func (h Handler) UpdateStudent(c *gin.Context) {

	student := models.Student{}

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, "error while validating studentId", http.StatusBadRequest, err.Error())
		return
	}
	student.Id = id

	if err := c.ShouldBindJSON(&student); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Store.StudentStorage().Update(student)
	if err != nil {
		handleResponse(c, "error while updating student", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Updated successfully", http.StatusOK, id)
}


func (h Handler) DeleteStudent(c *gin.Context) {

	student := models.Student{}

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, "error while validating studentId", http.StatusBadRequest, err.Error())
		return
	}
	student.Id = id

	if err := c.ShouldBindJSON(&student); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Store.StudentStorage().Delete(student)
	if err != nil {
		handleResponse(c, "error while deleting student", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Deleted successfully", http.StatusOK, id)
}



func (h Handler) GetStudent(c *gin.Context) {

	student := models.GetStudent{}

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, "error while validating studentId", http.StatusBadRequest, err.Error())
		return
	}
	student.Id = id

	if err := c.ShouldBindJSON(&student); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	std, err := h.Store.StudentStorage().GetStudent(student)
	if err != nil {
		handleResponse(c, "error while getting student", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Got successfully", http.StatusOK, std)
}

func (h Handler) GetAllStudents(c *gin.Context) {
	search := c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.Store.StudentStorage().GetAll(models.GetAllStudentsRequest{
		Search: search,
		Page:   page,
		Limit:  limit,
	})
	if err != nil {
		handleResponse(c, "error while getting all students", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "request successful", http.StatusOK, resp)
}
