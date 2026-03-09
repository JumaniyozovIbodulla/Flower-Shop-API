package models

import (
	"github.com/google/uuid"
)

type Order struct {
	ID          uuid.UUID `json:"id" example:"1b2d660e-c0dc-4da7-992f-c324acb26abd"`
	OrderNumber uint64    `json:"order_number" example:"79197919"`
	CustomerID  uuid.UUID `json:"customer_id" example:"9c1242b2-b211-4a91-ba6c-a58e903327fd"`
	TotalPrice  uint64    `json:"total_price" example:"25800"`
	Status      string    `json:"status" enums:"pending, confirmed, delivered, cancelled"`
	CreatedAt   int64     `json:"created_at"`
	UpdatedAt   int64     `json:"updated_at"`
}

type AddOrder struct {
	CustomerID uuid.UUID `json:"customer_id" example:"9c1242b2-b211-4a91-ba6c-a58e903327fd"`
	TotalPrice uint64    `json:"total_price" example:"25800"`
}
