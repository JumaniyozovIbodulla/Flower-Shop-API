package models

import (
	"github.com/google/uuid"
)

type Flower struct {
	ID          uuid.UUID `json:"id" example:"c9364df3-e183-4fe9-b53f-ba10dd8fc20d"`
	Title       string    `json:"title" example:"Wonderful life"`
	Description string    `json:"description" example:"Wonderful life is a classical bouquet made of beautiful autumn flowers and tender orchids."`
	UnitPrice   uint64    `json:"unit_price" example:"8600"`
	Stock       uint      `json:"stock" example:"1000"`
	IsActive    bool      `json:"is_active" example:"true"`
	CreatedAt   int64     `json:"created_at"`
}

type AddFlower struct {
	Title       string `json:"title" example:"Wonderful life"`
	Description string `json:"description" example:"Wonderful life is a classical bouquet made of beautiful autumn flowers and tender orchids."`
	UnitPrice   uint64 `json:"unit_price" example:"8600"`
	Stock       uint   `json:"stock" example:"1000"`
	IsActive    bool   `json:"is_active" example:"true"`
}

type UpdateFlower struct {
	ID          uuid.UUID `json:"id" example:"c9364df3-e183-4fe9-b53f-ba10dd8fc20d"`
	Title       string    `json:"title" example:"Wonderful life"`
	Description string    `json:"description" example:"Wonderful life is a classical bouquet made of beautiful autumn flowers and tender orchids."`
	UnitPrice   uint64    `json:"unit_price" example:"8600"`
	Stock       uint      `json:"stock" example:"1000"`
	IsActive    bool      `json:"is_active" example:"true"`
}

type GetAllFlowersRequest struct {
	SearchByTitle string `json:"search_by_title" example:"Wonderful life"`
	Page          uint64 `json:"page" example:"1"`
	Limit         uint64 `json:"limit" example:"10"`
}

type GetAllFlowersResponse struct {
	Flowers []Flower `json:"flowers"`
	Count   uint64   `json:"count"`
}
