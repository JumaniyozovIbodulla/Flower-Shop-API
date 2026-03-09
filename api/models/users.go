package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id" example:"9c1242b2-b211-4a91-ba6c-a58e903327fd"`
	Name         string    `json:"name" example:"Jumaniyozov Ibodulla"`
	Email        string    `json:"email" example:"jumaniyozovibodulla07@gmail.com"`
	Language     string    `json:"language" enums:"uz, ru, en"`
	PasswordHash string    `json:"password_hash" example:"PASSWORD_HASH"`
	CreatedAt    int64     `json:"created_at"`
}

type GetUser struct {
	ID        uuid.UUID `json:"id" example:"9c1242b2-b211-4a91-ba6c-a58e903327fd"`
	Name      string    `json:"name" example:"Jumaniyozov Ibodulla"`
	Email     string    `json:"email" example:"jumaniyozovibodulla07@gmail.com"`
	Language  string    `json:"language" enums:"uz, ru, en"`
	CreatedAt int64     `json:"created_at"`
}

type AddUser struct {
	Name         string `json:"name" example:"Jumaniyozov Ibodulla"`
	Email        string `json:"email" example:"jumaniyozovibodulla07@gmail.com"`
	Language     string `json:"language" enums:"uz, ru, en"`
	PasswordHash string `json:"password_hash" example:"PASSWORD_HASH"`
}

type UpdateUser struct {
	ID       uuid.UUID `json:"id" example:"9c1242b2-b211-4a91-ba6c-a58e903327fd"`
	Name     string    `json:"name" example:"Jumaniyozov Ibodulla"`
	Email    string    `json:"email" example:"jumaniyozovibodulla07@gmail.com"`
	Language string    `json:"language" enums:"uz, ru, en"`
}

type UpdateUserPassword struct {
	ID           uuid.UUID `json:"id" example:"9c1242b2-b211-4a91-ba6c-a58e903327fd"`
	PasswordHash string    `json:"password_hash" example:"PASSWORD_HASH"`
}

type GetAllUsersRequest struct {
	SearchByName  string `json:"search_by_name"`
	SearchByEmail string `json:"search_by_email"`
	Page          uint64 `json:"page"`
	Limit         uint64 `json:"limit"`
}

type GetAllUsersResponse struct {
	Users []GetUser `json:"users"`
	Count uint64    `json:"count"`
}
