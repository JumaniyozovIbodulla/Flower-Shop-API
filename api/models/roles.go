package models

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID          uuid.UUID `json:"id" example:"88361fd2-b050-4f41-ba25-e262aa7a3113"`
	Name        string    `json:"name" example:"Manager"`
	Description string    `json:"description" example:"Manages only orders"`
	CreatedAt   time.Time `json:"created_at"`
}

type AddRole struct {
	Name        string `json:"name" example:"Manager"`
	Description string `json:"description" example:"Manages only orders"`
}

type UpdateRole struct {
	ID          uuid.UUID `json:"id" example:"88361fd2-b050-4f41-ba25-e262aa7a3113"`
	Name        string    `json:"name" example:"Manager"`
	Description string    `json:"description" example:"Manages only orders"`
}

type GetAllRolesRequest struct {
	SearchByName string `json:"search_by_name" example:"Manager"`
	Page         uint64 `json:"page" example:"1"`
	Limit        uint64 `json:"limit" example:"10"`
}

type GetAllRolesResponse struct {
	Roles []Role `json:"roles"`
	Count uint64 `json:"count"`
}
