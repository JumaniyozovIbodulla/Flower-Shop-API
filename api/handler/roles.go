package handler

import (
	"errors"
	_ "flower-shop/api/docs"
	"flower-shop/api/models"
	"flower-shop/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

// CreateRole godoc
// @Security ApiKeyAuth
// @Router		/api/v1/role [POST]
// @Summary		CREATE A ROLE
// @Description	THIS API CREATE AN ROLE
// @Tags		ROLES
// @Accept		json
// @Produce		json
// @Param		roles body models.AddRole true "ROLE INFO"
// @Success		200  {object}  models.SeccessRequest
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateRole(c *gin.Context) {
	role := models.AddRole{}

	if err := c.ShouldBindJSON(&role); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.Role().Create(c.Request.Context(), role)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) &&
			pgErr.Code == "23505" &&
			pgErr.ConstraintName == "roles_name_key" {

			handleResponse(c, h.Log, "role already exists", http.StatusBadRequest, "role already exists")
			return
		}

		handleResponse(c, h.Log, "error while creating role", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseGet(c, h.Log, http.StatusCreated, models.SeccessRequest{
		Code:    "201",
		Message: "added new role",
	})
}

// GetAllRoles godoc
// @Security ApiKeyAuth
// @Router		/api/v1/roles [GET]
// @Summary		GET ALL ROLES
// @Description THIS API GET ALL ROLES
// @Tags		ROLES
// @Accept		json
// @Produce		json
// @Param		name  query string false "Search BY NAME"
// @Param		page query  int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.GetAllUsersResponse
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllRoles(c *gin.Context) {
	searchByName := c.Query("name")

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

	resp, err := h.Service.Role().GetAll(c.Request.Context(), models.GetAllRolesRequest{
		SearchByName: searchByName,
		Page:         page,
		Limit:        limit,
	})

	if err != nil {
		handleResponse(c, h.Log, "error while getting all roles", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseGet(c, h.Log, http.StatusOK, resp)
}

// UpdateRole godoc
// @Security ApiKeyAuth
// @Router		/api/v1/role [PUT]
// @Summary		UPDATE THE ROLE
// @Description	THIS API UPDATE THE ROLE
// @Tags		ROLES
// @Accept		json
// @Produce		json
// @Param		roles body models.UpdateRole true "ROLE INFO"
// @Success		200  {object}  models.SeccessRequest
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateRole(c *gin.Context) {
	var role models.UpdateRole

	if err := c.ShouldBindJSON(&role); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.Role().Update(c.Request.Context(), role)
	if err != nil {
		if errors.Is(err, pkg.RoleNotFoundErr) {
			handleResponse(c, h.Log, "error while deleting, role not found", http.StatusNotFound, err.Error())
			return
		}

		handleResponse(c, h.Log, "error while deleting role", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseGet(c, h.Log, http.StatusOK, models.SeccessRequest{
		Code:    "200",
		Message: "role updated successfully",
	})
}

// DeleteRole godoc
// @Security ApiKeyAuth
// @Router		/api/v1/role/{id} [DELETE]
// @Summary		DELETE A ROLE
// @Description	THIS API DELETE THE ROLE
// @Tags		ROLES
// @Accept		json
// @Produce		json
// @Param		id path string true "ROLE ID"
// @Success		200  {object}  models.SeccessRequest
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.Log, "error while validating user's id", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.Role().Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, pkg.UserNotFoundErr) {
			handleResponse(c, h.Log, "error while deleting, role not found", http.StatusNotFound, err.Error())
			return
		}

		handleResponse(c, h.Log, "error while deleting role", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseGet(c, h.Log, http.StatusOK, models.SeccessRequest{
		Code:    "200",
		Message: "role deleted successfully",
	})
}
