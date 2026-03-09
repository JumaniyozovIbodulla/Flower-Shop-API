package handler

import (
	"errors"
	_ "flower-shop/api/docs"
	"flower-shop/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

// CreateUser godoc
// @Security ApiKeyAuth
// @Router		/api/v1/user [POST]
// @Summary		CREATE AN USER
// @Description	THIS API CREATE AN USER
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
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) &&
			pgErr.Code == "23505" &&
			pgErr.ConstraintName == "users_email_key" {
			
			handleResponse(c, h.Log, "email already exists", http.StatusBadRequest, "email already exists")
			return
		}

		handleResponse(c, h.Log, "error while creating user", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseGet(c, h.Log, http.StatusCreated, models.SeccessRequest{
		Code:    "201",
		Message: "Yangi foydalanuvchi qo'shildi",
	})
}

// GetAllUsers godoc
// @Security ApiKeyAuth
// @Router		/api/v1/users [GET]
// @Summary		GET ALL USERS
// @Description THIS API GET ALL USERS
// @Tags		USERS
// @Accept		json
// @Produce		json
// @Param		search_by_name  query string false "Search BY NAME"
// @Param		search_by_email query string false "Search BY EMAIL"
// @Param		page query  int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.GetAllUsersResponse
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllUsers(c *gin.Context) {
	searchByName := c.Query("search_by_name")
	searchByEmail := c.Query("search_by_email")

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
		SearchByName:  searchByName,
		SearchByEmail: searchByEmail,
		Page:          page,
		Limit:         limit,
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
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.Log, "error while validating user's id", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.User().Delete(c.Request.Context(), id)
	if err != nil {
		handleResponse(c, h.Log, "error while deleting user", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseGet(c, h.Log, http.StatusOK, models.SeccessRequest{
		Code:    "200",
		Message: "User deleted successfully",
	})
}
