package models

import (
	"github.com/google/uuid"
)

type Cart struct {
	ID         uuid.UUID `json:"id" example:"0b754271-a695-4526-a173-d693ec2d4c12"`
	CustomerID uuid.UUID `json:"customer_id" example:"9c1242b2-b211-4a91-ba6c-a58e903327fd"`
	CreatedAt  int64     `json:"created_at"`
}

type AddCart struct {
	CustomerID uuid.UUID `json:"customer_id" example:"9c1242b2-b211-4a91-ba6c-a58e903327fd"`
}

type GetAllCartsRequest struct {
	SearchByCustomerID uuid.UUID `json:"search_by_customer_id" example:"0b754271-a695-4526-a173-d693ec2d4c12"`
	Page               uint64    `json:"page" example:"1"`
	Limit              uint64    `json:"limit" example:"10"`
}
