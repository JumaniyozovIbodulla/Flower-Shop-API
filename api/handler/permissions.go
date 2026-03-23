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

// CreatePermission godoc
// @Security ApiKeyAuth
// @Router		/api/v1/permission [POST]
// @Summary		CREATE A PERMISSION
// @Description	THIS API CREATE AN PERMISSION
// @Tags		PERMISSIONS
// @Accept		json
// @Produce		json
// @Param		permissions body models.AddPermission true "PERMISSION INFO"
// @Success		200  {object}  models.SeccessRequest
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreatePermission(c *gin.Context) {
	permission := models.AddPermission{}

	if err := c.ShouldBindJSON(&permission); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.Permission().Create(c.Request.Context(), permission)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) &&
			pgErr.Code == "23505" &&
			pgErr.ConstraintName == "permissions_name_key" {

			handleResponse(c, h.Log, "permission already exists", http.StatusBadRequest, "permission already exists")
			return
		}

		handleResponse(c, h.Log, "error while creating permission", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseGet(c, h.Log, http.StatusCreated, models.SeccessRequest{
		Code:    "201",
		Message: "added new permission",
	})
}

// GetAllPermissions godoc
// @Security ApiKeyAuth
// @Router		/api/v1/permissions [GET]
// @Summary		GET ALL PERMISSIONS
// @Description THIS API GET ALL PERMISSIONS
// @Tags		PERMISSIONS
// @Accept		json
// @Produce		json
// @Param		name  query string false "Search BY NAME"
// @Param		page  query  int   false "page"
// @Param		limit query  int    false "limit"
// @Success		200  {object}  models.GetAllUsersResponse
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllPermissions(c *gin.Context) {
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

	resp, err := h.Service.Permission().GetAll(c.Request.Context(), models.GetAllPermissionsRequest{
		SearchByName: searchByName,
		Page:         page,
		Limit:        limit,
	})

	if err != nil {
		handleResponse(c, h.Log, "error while getting all permissions", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseGet(c, h.Log, http.StatusOK, resp)
}

// UpdatePermission godoc
// @Security ApiKeyAuth
// @Router		/api/v1/permission [PUT]
// @Summary		UPDATE THE PERMISSION
// @Description	THIS API UPDATE THE PERMISSION
// @Tags		PERMISSIONS
// @Accept		json
// @Produce		json
// @Param		permission body models.UpdatePermission true "ROLE INFO"
// @Success		200  {object}  models.SeccessRequest
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdatePermission(c *gin.Context) {
	var permission models.UpdatePermission

	if err := c.ShouldBindJSON(&permission); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.Permission().Update(c.Request.Context(), permission)
	if err != nil {
		if errors.Is(err, pkg.NotFoundErr) {
			handleResponse(c, h.Log, "error while deleting, permission not found", http.StatusNotFound, err.Error())
			return
		}

		handleResponse(c, h.Log, "error while deleting permission", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseGet(c, h.Log, http.StatusOK, models.SeccessRequest{
		Code:    "200",
		Message: "permission updated successfully",
	})
}

// DeletePermission godoc
// @Security ApiKeyAuth
// @Router		/api/v1/permission/{id} [DELETE]
// @Summary		DELETE A PERMISSION
// @Description	THIS API DELETE THE PERMISSION
// @Tags		PERMISSIONS
// @Accept		json
// @Produce		json
// @Param		id path string true "PERMISSION ID"
// @Success		200  {object}  models.SeccessRequest
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) DeletePermission(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.Log, "error while validating permission's id", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.Permission().Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, pkg.NotFoundErr) {
			handleResponse(c, h.Log, "error while deleting, permission not found", http.StatusNotFound, err.Error())
			return
		}

		handleResponse(c, h.Log, "error while deleting permission", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseGet(c, h.Log, http.StatusOK, models.SeccessRequest{
		Code:    "200",
		Message: "permission deleted successfully",
	})
}

