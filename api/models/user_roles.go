package models

import "github.com/google/uuid"

type UserRole struct {
	UserID uuid.UUID `json:"user_id" example:"9c1242b2-b211-4a91-ba6c-a58e903327fd"`
	RoleID uuid.UUID `json:"role_id" example:"88361fd2-b050-4f41-ba25-e262aa7a3113"`
}

