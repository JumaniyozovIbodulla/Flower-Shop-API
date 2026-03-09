package models

import "github.com/google/uuid"

type OrderItem struct {
	ID        uuid.UUID `json:"id" example:"a2edfa3a-877c-4c69-b16b-e6313a3ce8c5"`
	OrderID   uuid.UUID `json:"order_id" example:"1b2d660e-c0dc-4da7-992f-c324acb26abd"`
	FlowerID  uuid.UUID `json:"flower_id" example:"c9364df3-e183-4fe9-b53f-ba10dd8fc20d"`
	Quantity  uint      `json:"quantity" example:"3"`
	UnitPrice uint64    `json:"unit_price" example:"25800"`
}

type AddOrderItem struct {
	OrderID   uuid.UUID `json:"order_id" example:"1b2d660e-c0dc-4da7-992f-c324acb26abd"`
	FlowerID  uuid.UUID `json:"flower_id" example:"c9364df3-e183-4fe9-b53f-ba10dd8fc20d"`
	Quantity  uint      `json:"quantity" example:"3"`
	UnitPrice uint64    `json:"unit_price" example:"25800"`
}

type GetAllOrderItemsRequest struct {
	FilterByOrderID uuid.UUID `json:"filter_by_order_id"`
	Page            uint64    `json:"page" example:"1"`
	Limit           uint64    `json:"limit" example:"10"`
}

type GetAllOrderItemsResponse struct {
	OrderItems []OrderItem `json:"order_items"`
	Count      uint64      `json:"count"`
}

