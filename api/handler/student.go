package handler

import (
	_ "backend_course/lms/api/docs"
	"backend_course/lms/api/models"
	"backend_course/lms/pkg/check"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Router		/student [POST]
// @Summary		create a student
// @Description	This api create a student and returns its id
// @Tags		student
// @Accept		json
// @Produce		json
// @Param		student body models.AddStudent true "student"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateStudent(c *gin.Context) {
	student := models.AddStudent{}

	if err := c.ShouldBindJSON(&student); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	if err := check.ValidateYear(student.Age); err != nil {
		handleResponse(c, h.Log, "error while validating student age, year: "+strconv.Itoa(student.Age), http.StatusBadRequest, err.Error())

		return
	}

	id, err := h.Service.Student().Create(c.Request.Context(), student)
	if err != nil {
		handleResponse(c, h.Log, "error while creating student", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Created successfully", http.StatusOK, id)
}

// @Router		/student/{id} [PUT]
// @Summary		update a student
// @Description	This api update a student and returns its id
// @Tags		student
// @Accept		json
// @Produce		json
// @Param		student body models.AddStudent true "student"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateStudent(c *gin.Context) {

	student := models.Student{}

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.Log, "error while validating studentId", http.StatusBadRequest, err.Error())
		return
	}
	student.Id = id

	if err := c.ShouldBindJSON(&student); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Service.Student().Update(c.Request.Context(), student)
	if err != nil {
		handleResponse(c, h.Log, "error while updating student", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Updated successfully", http.StatusOK, id)
}

// @Router		/student/{id} [PATCH]
// @Summary		update a student's status
// @Description	This api update a student's status and returns its id
// @Tags		student
// @Accept		json
// @Produce		json
// @Param		student body models.Student true "student"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateStudentStatus(c *gin.Context) {

	student := models.Student{}

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.Log, "error while validating studentId", http.StatusBadRequest, err.Error())
		return
	}
	student.Id = id

	if err := c.ShouldBindJSON(&student); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Service.Student().UpdateStatus(c.Request.Context(), student)
	if err != nil {
		handleResponse(c, h.Log, "error while updating student", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Updated successfully", http.StatusOK, id)
}

// @Router		/student/{id} [DELETE]
// @Summary		delete a student
// @Description	This api delete a student
// @Tags		student
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.Log, "error while validating studentId", http.StatusBadRequest, err.Error())
		return
	}
	if err := h.Service.Student().Delete(c.Request.Context(), id); err != nil {
		handleResponse(c, h.Log, "error while deleting student", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Deleted successfully", http.StatusOK, id)
}

// @Router		/student/{id} [GET]
// @Summary		get a student
// @Description	This api get a student
// @Tags		student
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetStudent(c *gin.Context) {

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.Log, "error while validating studentId", http.StatusBadRequest, err.Error())
		return
	}

	std, err := h.Service.Student().GetStudent(c.Request.Context(), id)
	if err != nil {
		handleResponse(c, h.Log, "error while getting student", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Got successfully", http.StatusOK, std)
}

// @Router		/students [GET]
// @Summary		get  students
// @Description	This api get all students
// @Tags		student
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllStudents(c *gin.Context) {
	search := c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c, h.Log, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, h.Log, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.Service.Student().GetAll(c.Request.Context(), models.GetAllStudentsRequest{
		Search: search,
		Page:   page,
		Limit:  limit,
	})
	if err != nil {
		handleResponse(c, h.Log, "error while getting all students", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, h.Log, "request successful", http.StatusOK, resp)
}
