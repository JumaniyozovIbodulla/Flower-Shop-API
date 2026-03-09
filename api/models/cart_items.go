package models

import "github.com/google/uuid"

type CartItem struct {
	ID        uuid.UUID `json:"id" example:"c735154c-ebb9-432b-aa49-2821c3e5411e"`
	CartID    uuid.UUID `json:"cart_id" example:"0b754271-a695-4526-a173-d693ec2d4c12"`
	FlowerID  uuid.UUID `json:"flower_id" example:"c9364df3-e183-4fe9-b53f-ba10dd8fc20d"`
	Quantity  uint      `json:"quantity" example:"3"`
	UnitPrice uint64    `json:"unit_price" example:""`
}

type AddCartItem struct {
	CartID    uuid.UUID `json:"cart_id" example:"0b754271-a695-4526-a173-d693ec2d4c12"`
	FlowerID  uuid.UUID `json:"flower_id" example:"c9364df3-e183-4fe9-b53f-ba10dd8fc20d"`
	Quantity  uint      `json:"quantity" example:"3"`
	UnitPrice uint64    `json:"unit_price" example:"8600"`
}

type GetAllCartItemsRequest struct {
	SearchByCartID uuid.UUID `json:"search_by_cart_id" example:"0b754271-a695-4526-a173-d693ec2d4c12"`
	Page           uint64    `json:"page" example:"1"`
	Limit          uint64    `json:"limit" example:"10"`
}

type GetAllCartItemsResponse struct {
	CartItems []CartItem `json:"cart_items"`
	Count     uint64     `json:"count"`
}

