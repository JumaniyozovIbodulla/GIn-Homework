package handler

import (
	_ "backend_course/lms/api/docs"
	"backend_course/lms/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Router		/subject [POST]
// @Summary		create a subject
// @Description	This api create a subject and returns its id
// @Tags		subject
// @Accept		json
// @Produce		json
// @Param		subject body models.AddSubject true "subject"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateSubject(c *gin.Context) {
	subject := models.AddSubject{}

	if err := c.ShouldBindJSON(&subject); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.Subjects().Create(c.Request.Context(), subject)
	if err != nil {
		handleResponse(c, h.Log, "error while creating student", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Created successfully", http.StatusOK, id)
}

// @Router		/subject/{id} [PUT]
// @Summary		update a subject
// @Description	This api update a subject and returns its id
// @Tags		subject
// @Accept		json
// @Produce		json
// @Param		subject body models.Subjects true "subject"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateSubject(c *gin.Context) {

	subject := models.Subjects{}

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.Log, "error while validating studentId", http.StatusBadRequest, err.Error())
		return
	}
	subject.Id = id

	if err := c.ShouldBindJSON(&subject); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Service.Subjects().Update(c.Request.Context(), subject)
	if err != nil {
		handleResponse(c, h.Log, "error while updating subject", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Updated successfully", http.StatusOK, id)
}


// @Router		/subject/{id} [DELETE]
// @Summary		delete a subject
// @Description	This api delete a subject
// @Tags		subject
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) DeleteSubject(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.Log, "error while validating studentId", http.StatusBadRequest, err.Error())
		return
	}
	if err := h.Service.Subjects().Delete(c.Request.Context(), id); err != nil {
		handleResponse(c, h.Log, "error while deleting student:", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Deleted successfully", http.StatusOK, id)
}

// @Router		/subject/{id} [GET]
// @Summary		get a subject
// @Description	This api get a subject
// @Tags		subject
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetSubject(c *gin.Context) {

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.Log, "error while validating subjectId", http.StatusBadRequest, err.Error())
		return
	}

	std, err := h.Service.Subjects().GetSubject(c.Request.Context(), id)
	if err != nil {
		handleResponse(c, h.Log, "error while getting subject", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Got successfully", http.StatusOK, std)
}

// @Router		/subjects [GET]
// @Summary		get  subjects
// @Description	This api get all subjects
// @Tags		subject
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllSubjects(c *gin.Context) {
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

	resp, err := h.Service.Subjects().GetAll(c.Request.Context(), models.GetAllSubjectsRequest{
		Search: search,
		Page:   page,
		Limit:  limit,
	})
	if err != nil {
		handleResponse(c, h.Log, "error while getting all subjects", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, h.Log, "request successful", http.StatusOK, resp)
}
