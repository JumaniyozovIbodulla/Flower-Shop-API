package models

import (
	"github.com/google/uuid"
)

type OrderStatusHistory struct {
	ID        uuid.UUID `json:"id" example:"6c8122cb-1ad8-4b8f-81d2-ecfecd150bab"`
	OrderID   uuid.UUID `json:"order_id" example:"1b2d660e-c0dc-4da7-992f-c324acb26abd"`
	Status    string    `json:"status" enums:"pending, confirmed, delivered, cancelled"`
	CreatedAt int64     `json:"created_at"`
}

type AddOrderStatusHistory struct {
	OrderID uuid.UUID `json:"order_id" example:"1b2d660e-c0dc-4da7-992f-c324acb26abd"`
	Status  string    `json:"status" enums:"pending, confirmed, delivered, cancelled"`
}

type GetAllOrderStatusHistoryRequest struct {
	SearchByOrderID uuid.UUID `json:"searchByOrderID"`
	Page            uint64    `json:"page" example:"1"`
	Limit           uint64    `json:"limit" example:"10"`
}

type GetAllOrderStatusHistoryResponse struct {
	OrderStatusHistories []OrderStatusHistory `json:"order_status_histories"`
	Count                uint64               `json:"count"`
}
