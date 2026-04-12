package handler

import (
	"errors"
	_ "flower-shop/api/docs"
	"flower-shop/api/models"
	"flower-shop/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

// CreateRolePermission godoc
// @Security ApiKeyAuth
// @Router		/api/v1/role-permission [POST]
// @Summary		CREATE A ROLE PERMISSION
// @Description	THIS API CREATE A ROLE PERMISSION
// @Tags		ROLE-PERMISSIONS
// @Accept		json
// @Produce		json
// @Param		role_permissions body models.RolePermission true "ROLE PERMISSION"
// @Success		200  {object}  models.SeccessRequest
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateRolePermission(c *gin.Context) {
	rolePermission := models.RolePermission{}

	if err := c.ShouldBindJSON(&rolePermission); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.RolePermissions().Create(c.Request.Context(), rolePermission)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) &&
			pgErr.Code == "23505" &&
			pgErr.ConstraintName == "permissions_name_key" {

			handleResponse(c, h.Log, "permission already attached", http.StatusBadRequest, "permission already attached")
			return
		}

		handleResponse(c, h.Log, "error while attaching permission to the role", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseGet(c, h.Log, http.StatusCreated, models.SeccessRequest{
		Code:    "201",
		Message: "attached permission to role",
	})
}

// DeleteRolePermission godoc
// @Security ApiKeyAuth
// @Router		/api/v1/role-permission [DELETE]
// @Summary		DELETE THE ROLE PERMISSION
// @Description	THIS API DELETES THE ROLE PERMISSION
// @Tags		ROLE-PERMISSIONS
// @Accept		json
// @Produce		json
// @Param		rolePermission body models.RolePermission true "ROLE PERMISSION"
// @Success		200  {object}  models.SeccessRequest
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) DeleteRolePermission(c *gin.Context) {
	var req models.RolePermission

	if err := c.ShouldBindJSON(&req); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.RolePermissions().Delete(c.Request.Context(), req)
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
		Message: "the permission removed from the role successfully",
	})
}
