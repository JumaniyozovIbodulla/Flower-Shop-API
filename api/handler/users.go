package handler

import (
	_ "flower-shop/api/docs"
	"flower-shop/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// CreateUser godoc
// @Security ApiKeyAuth
// @Router		/api/v1/user [POST]
// @Summary		CREATE AN USER
// @Description	This api CREATE AN USER
// @Tags		USERS
// @Accept		json
// @Produce		json
// @Param		users body models.AddUser true "USER INFO"
// @Success		200  {object}  models.SeccessRequest
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateUser(c *gin.Context) {
	user := models.AddUser{}

	if err := c.ShouldBindJSON(&user); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.User().Create(c.Request.Context(), user)
	if err != nil {
		handleResponse(c, h.Log, "error while creating user", http.StatusBadRequest, err.Error())
		return
	}

	handleResponseGet(c, h.Log, http.StatusCreated, models.SeccessRequest{
		Code:    "201",
		Message: "Yangi foydalanuvchi qo'shildi",
	})
}

// GetAllUser godoc
// @Security ApiKeyAuth
// @Router		/api/v1/users [GET]
// @Summary		GET ALL USERS
// @Description THIS API GET ALL USERS
// @Tags		USERS
// @Accept		json
// @Produce		json
// @Param		search_by_full_name query string false "Search BY FULL NAME"
// @Param		search_by_username query string false "Search BY USERNAME"
// @Param		search_by_id query int false "Search BY ID"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.GetAllUsersResponse
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllUser(c *gin.Context) {
	searchByFullName := c.Query("search_by_full_name")
	searchByUsername := c.Query("search_by_username")
	searchByID := c.Query("search_by_id")

	if searchByID == "" {
		searchByID = "0"
	}

	userID := cast.ToInt64(searchByID)

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

	resp, err := h.Service.User().GetAll(c.Request.Context(), models.GetAllUsersRequest{
		SearchByFullName: searchByFullName,
		SearchByUsername: searchByUsername,
		SearchByID:       userID,
		Page:             page,
		Limit:            limit,
	})

	if err != nil {
		handleResponse(c, h.Log, "error while getting all users", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseGet(c, h.Log, http.StatusOK, resp)
}

// DeleteUser godoc
// @Security ApiKeyAuth
// @Router		/api/v1/user/{id} [DELETE]
// @Summary		DELETE AN USER
// @Description	THIS API DELETE AN USER
// @Tags		USERS
// @Accept		json
// @Produce		json
// @Param		id path string true "USER ID"
// @Success		200  {object}  models.SeccessRequest
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	userID := cast.ToInt64(id)

	err := h.Service.User().Delete(c.Request.Context(), userID)
	if err != nil {
		handleResponse(c, h.Log, "error while deleting user", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseGet(c, h.Log, http.StatusOK, models.SeccessRequest{
		Code:    "200",
		Message: "user deleted successfully",
	})
}
