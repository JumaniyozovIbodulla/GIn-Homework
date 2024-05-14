package handler

import (
	_ "backend_course/lms/api/docs"
	"backend_course/lms/api/models"
	"backend_course/lms/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Security ApiKeyAuth
// @Router		/teacher [POST]
// @Summary		create a teacher
// @Description	This api create a teacher and returns its id
// @Tags		teacher
// @Accept		json
// @Produce		json
// @Param		teacher body models.AddTeacher true "teacher"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateTeacher(c *gin.Context) {
	teacher := models.AddTeacher{}

	if err := c.ShouldBindJSON(&teacher); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	password, err := pkg.HashPassword(teacher.Password)

	if err != nil {
		handleResponse(c, h.Log, "error while hashing password", http.StatusBadRequest, err.Error())
		return
	}

	teacher.Password = password

	id, err := h.Service.Teacher().Create(c.Request.Context(), teacher)
	if err != nil {
		handleResponse(c, h.Log, "error while creating teacher", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Created successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// @Security	ApiKeyAuth
// @Router		/teacher/{id} [PUT]
// @Summary		update a teacher
// @Description	This api update a teacher and returns its id
// @Tags		teacher
// @Accept		json
// @Produce		json
// @Param		teacher body models.AddTeacher true "teacher"
// @Param		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateTeacher(c *gin.Context) {

	teacher := models.Teacher{}

	_, err := getAuthInfo(c)
	if err != nil {
		handleResponse(c, h.Log, "unauthorized", http.StatusUnauthorized, err.Error())
		return
	}

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.Log, "error while validating teacherId", http.StatusBadRequest, err.Error())
		return
	}
	teacher.Id = id

	if err := c.ShouldBindJSON(&teacher); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	id, err = h.Service.Teacher().Update(c.Request.Context(), teacher)
	if err != nil {
		handleResponse(c, h.Log, "error while updating teacher", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Updated successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// @Router		/teacher/{id} [DELETE]
// @Summary		delete a teacher
// @Description	This api delete a teacher
// @Tags		teacher
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) DeleteTeacher(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.Log, "error while validating teacherId", http.StatusBadRequest, err.Error())
		return
	}
	if err := h.Service.Teacher().Delete(c.Request.Context(), id); err != nil {
		handleResponse(c, h.Log, "error while deleting teacher", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Deleted successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// @Router		/teacher/{id} [GET]
// @Summary		get a teacher
// @Description	This api get a teacher
// @Tags		teacher
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetTeacher(c *gin.Context) {

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.Log, "error while validating teacherId", http.StatusBadRequest, err.Error())
		return
	}

	std, err := h.Service.Teacher().GetTeacher(c.Request.Context(), id)
	if err != nil {
		handleResponse(c, h.Log, "error while getting teacher", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Got successfully", http.StatusOK, std)
}

// @Security ApiKeyAuth
// @Router		/teachers [GET]
// @Summary		get  all teachers
// @Description	This api get all teachers
// @Tags		teacher
// @Accept		json
// @Produce		json
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllTeachers(c *gin.Context) {
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

	resp, err := h.Service.Teacher().GetAll(c.Request.Context(), models.GetAllTeachersRequest{
		Limit: limit,
		Page: page,
		Search: search,
	})
	if err != nil {
		handleResponse(c, h.Log, "error while getting all teachers", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, h.Log, "request successful", http.StatusOK, resp)
}


// @Security ApiKeyAuth
// @Router		/check-teacher/{id} [GET]
// @Summary		get a teacher's lesson
// @Description	This api get a teacher's lesson
// @Tags		teacher
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetTeacherLesson(c *gin.Context) {

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.Log, "error while validating teacherId", http.StatusBadRequest, err.Error())
		return
	}

	std, err := h.Service.Teacher().CheckTeacherLesson(c.Request.Context(), id)
	if err != nil {
		handleResponse(c, h.Log, "error while getting check teacher's lesson", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Got successfully", http.StatusOK, std)
}