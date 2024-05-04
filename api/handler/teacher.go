package handler

import (
	_ "backend_course/lms/api/docs"
	"backend_course/lms/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Router		/teacher [POST]
// @Summary		create a teacher
// @Description	This api create a teacher and returns its id
// @Tags		teacher
// @Accept		json
// @Produce		json
// @Param		teacher body models.Teacher true "teacher"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateTeacher(c *gin.Context) {
	teacher := models.Teacher{}

	if err := c.ShouldBindJSON(&teacher); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Store.TeacherStorage().Create(teacher)
	if err != nil {
		handleResponse(c, "error while creating teacher", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Created successfully", http.StatusOK, id)
}

// @Router		/teacher/{id} [PUT]
// @Summary		update a teacher
// @Description	This api update a teacher and returns its id
// @Tags		teacher
// @Accept		json
// @Produce		json
// @Param		student body models.Teacher true "teacher"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateTeacher(c *gin.Context) {

	teacher := models.Teacher{}

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, "error while validating teacherId", http.StatusBadRequest, err.Error())
		return
	}
	teacher.Id = id

	if err := c.ShouldBindJSON(&teacher); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Store.TeacherStorage().Update(teacher)
	if err != nil {
		handleResponse(c, "error while updating teacher", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Updated successfully", http.StatusOK, id)
}

// @Router		/teacher/{id} [DELETE]
// @Summary		delete a teacher
// @Description	This api delete a teacher
// @Tags		teacher
// @Accept		json
// @Produce		json
// @Param		teacher body models.Teacher true "teacher"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) DeleteTeacher(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, "error while validating teacherId", http.StatusBadRequest, err.Error())
		return
	}
	if err := h.Store.TeacherStorage().Delete(id); err != nil {
		handleResponse(c, "error while deleting teacher", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Deleted successfully", http.StatusOK, id)
}

// @Router		/teacher/{id} [GET]
// @Summary		Get a teacher
// @Description	This api get a teacher
// @Tags		teacher
// @Accept		json
// @Produce		json
// @Param		teacher body models.teacher true "teacher"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetTeacher(c *gin.Context) {

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, "error while validating teacherId", http.StatusBadRequest, err.Error())
		return
	}

	std, err := h.Store.TeacherStorage().GetTeacher(id)
	if err != nil {
		handleResponse(c, "error while getting teacher", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Got successfully", http.StatusOK, std)
}

// @Router		/teacher [GET]
// @Summary		Get  teacher
// @Description	This api get all teachers
// @Tags		teacher
// @Accept		json
// @Produce		json
// @Param		teacher body models.Teacher true "teacher"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllTeachers(c *gin.Context) {
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

	resp, err := h.Store.TeacherStorage().GetAll(models.GetAllTeachersRequest{
		Limit: limit,
		Page: page,
		Search: search,
	})
	if err != nil {
		handleResponse(c, "error while getting all teachers", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "request successful", http.StatusOK, resp)
}