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

// CreateUserRole godoc
// @Security ApiKeyAuth
// @Router		/api/v1/user-role [POST]
// @Summary		CREATE A USER-ROLE
// @Description	THIS API CREATE A USER-ROLE
// @Tags		USER-ROLES
// @Accept		json
// @Produce		json
// @Param		user_role body models.UserRole true "User Role model"
// @Success		200  {object}  models.SeccessRequest
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateUserRole(c *gin.Context) {
	userRole := models.UserRole{}

	if err := c.ShouldBindJSON(&userRole); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.UserRoles().Create(c.Request.Context(), userRole)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) &&
			pgErr.Code == "23505" &&
			pgErr.ConstraintName == "user_roles_pkey" {

			handleResponse(c, h.Log, "role already attached", http.StatusBadRequest, "role already attached")
			return
		}

		handleResponse(c, h.Log, "error while attaching role to the user", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseGet(c, h.Log, http.StatusCreated, models.SeccessRequest{
		Code:    "201",
		Message: "attached to the user",
	})
}

// DeleteUserRole godoc
// @Security ApiKeyAuth
// @Router		/api/v1/user-role [DELETE]
// @Summary		DELETE THE USER ROLE
// @Description	THIS API DELETES THE USER ROLE
// @Tags		USER-ROLES
// @Accept		json
// @Produce		json
// @Param		user_role body models.UserRole true "User Role"
// @Success		200  {object}  models.SeccessRequest
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) DeleteUserRole(c *gin.Context) {
	var req models.UserRole

	if err := c.ShouldBindJSON(&req); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.UserRoles().Delete(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, pkg.NotFoundErr) {
			handleResponse(c, h.Log, "error while deleting the role", http.StatusNotFound, err.Error())
			return
		}

		handleResponse(c, h.Log, "error while deleting the role", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseGet(c, h.Log, http.StatusOK, models.SeccessRequest{
		Code:    "200",
		Message: "the role removed successfully",
	})
}

// UpdateUserRole godoc
// @Security ApiKeyAuth
// @Router		/api/v1/user-role [PUT]
// @Summary		UPDATE THE USER ROLE
// @Description	THIS API UPDATES THE USER ROLE
// @Tags		USER-ROLES
// @Accept		json
// @Produce		json
// @Param		user_role body models.UserRole true "User Role"
// @Success		200  {object}  models.SeccessRequest
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateUserRole(c *gin.Context) {
	var req models.UserRole

	if err := c.ShouldBindJSON(&req); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.UserRoles().Update(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, pkg.NotFoundErr) {
			handleResponse(c, h.Log, "error while updating the role", http.StatusNotFound, err.Error())
			return
		}

		handleResponse(c, h.Log, "error while updating the role", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseGet(c, h.Log, http.StatusOK, models.SeccessRequest{
		Code:    "200",
		Message: "the role updated successfully",
	})
}
