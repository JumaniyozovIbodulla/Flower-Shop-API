package models

import "github.com/google/uuid"

type RolePermission struct {
	RoleID       uuid.UUID `json:"role_id" example:"88361fd2-b050-4f41-ba25-e262aa7a3113"`
	PermissionID uuid.UUID `json:"permission_id" example:"ccae8606-8be2-4621-9086-4144214b167b"`
}

