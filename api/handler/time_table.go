package handler

import (
	_ "backend_course/lms/api/docs"
	"backend_course/lms/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Router		/time [POST]
// @Summary		create a time table
// @Description	This api create a time table and returns its id
// @Tags		time_table
// @Accept		json
// @Produce		json
// @Param		time_table body models.Time true "time table"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateTime(c *gin.Context) {
	time := models.Time{}
	if err := c.ShouldBindJSON(&time); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Store.TimeStorage().Create(c.Request.Context(), time)
	if err != nil {
		handleResponse(c, "error while creating time table", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Created successfully", http.StatusOK, id)
}

// @Router		/time/{id} [PUT]
// @Summary		update a time table
// @Description	This api update a time table and returns its id
// @Tags		time_table
// @Accept		json
// @Produce		json
// @Param		time_table body models.Time true "time table"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateTime(c *gin.Context) {

	time := models.Time{}

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, "error while validating studentId", http.StatusBadRequest, err.Error())
		return
	}
	time.Id = id

	if err := c.ShouldBindJSON(&time); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Store.TimeStorage().Update(c.Request.Context(), time)
	if err != nil {
		handleResponse(c, "error while updating time table", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Updated successfully", http.StatusOK, id)
}


// @Router		/time/{id} [DELETE]
// @Summary		delete a time table
// @Description	This api delete a time table
// @Tags		time_table
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) DeleteTime(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, "error while validating timeId", http.StatusBadRequest, err.Error())
		return
	}
	if err := h.Store.TimeStorage().Delete(c.Request.Context(), id); err != nil {
		handleResponse(c, "error while deleting student", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Deleted successfully", http.StatusOK, id)
}

// @Router		/time/{id} [GET]
// @Summary		Get a time table
// @Description	This api get a time_table
// @Tags		time_table
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetTime(c *gin.Context) {

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, "error while validating timeId", http.StatusBadRequest, err.Error())
		return
	}

	std, err := h.Store.TimeStorage().GetTime(c.Request.Context(), id)
	if err != nil {
		handleResponse(c, "error while getting time table", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Got successfully", http.StatusOK, std)
}

// @Router		/time-tables [GET]
// @Summary		Get  time tables
// @Description	This api get all time tables
// @Tags		time_table
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllTimeTables(c *gin.Context) {
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

	resp, err := h.Store.TimeStorage().GetAll(c.Request.Context(), models.GetAllTimeRequest{
		Search: search,
		Page:   page,
		Limit:  limit,
	})
	if err != nil {
		handleResponse(c, "error while getting all time tables", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "request successful", http.StatusOK, resp)
}
