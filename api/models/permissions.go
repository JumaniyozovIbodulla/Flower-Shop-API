package models

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID          uuid.UUID `json:"id" example:"ccae8606-8be2-4621-9086-4144214b167b"`
	Name        string    `json:"name" example:"create_flower"`
	Description string    `json:"description" example:"Can add a new flowers"`
	CreatedAt   time.Time `json:"created_at"`
}

type AddPermission struct {
	Name        string `json:"name" example:"create_flower"`
	Description string `json:"description" example:"Can add a new flowers"`
}

type UpdatePermission struct {
	ID          uuid.UUID `json:"id" example:"ccae8606-8be2-4621-9086-4144214b167b"`
	Name        string    `json:"name" example:"delete_flower"`
	Description string    `json:"description" example:"Can remove a flower"`
}

type GetAllPermissionsRequest struct {
	SearchByName string `json:"search_by_name" example:"create_flower"`
	Page         uint64 `json:"page" example:"1"`
	Limit        uint64 `json:"limit" example:"10"`
}

type GetAllPermissionsResponse struct {
	Permissions []Permission `json:"permissions"`
	Count       uint64       `json:"count"`
}
