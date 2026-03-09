package models

import (
	"github.com/google/uuid"
)

type FlowerMedia struct {
	ID          uuid.UUID `json:"id" example:"a08e0613-ab09-4010-92c8-06c75bb5fbcb"`
	FlowerID    uuid.UUID `json:"flower_id" example:"c9364df3-e183-4fe9-b53f-ba10dd8fc20d"`
	ContentType string    `json:"content_type" enums:"image, video"`
	ObjectName  string    `json:"object_name" example:"flowers/123124wq.png"`
	CreatedAt   int64     `json:"created_at"`
}

type AddFlowerMedia struct {
	FlowerID    uuid.UUID `json:"flower_id" example:"c9364df3-e183-4fe9-b53f-ba10dd8fc20d"`
	ContentType string    `json:"content_type" enums:"image, video"`
	ObjectName  string    `json:"object_name" example:"flowers/123124wq.png" swaggerignore:"true"`
}

type UpdateFlowerMedia struct {
	ID          uuid.UUID `json:"id" example:"a08e0613-ab09-4010-92c8-06c75bb5fbcb"`
	FlowerID    uuid.UUID `json:"flower_id" example:"c9364df3-e183-4fe9-b53f-ba10dd8fc20d"`
	ContentType string    `json:"content_type" enums:"image, video"`
	ObjectName  string    `json:"object_name" example:"flowers/123124wq.png" swaggerignore:"true"`
}

type GetAllFlowerMediasRequest struct {
	SearchByFlowerID uuid.UUID `json:"searchByFlowerID" example:"c9364df3-e183-4fe9-b53f-ba10dd8fc20d"`
	Page             uint64    `json:"page" example:"1"`
	Limit            uint64    `json:"limit" example:"10"`
}

type GetAllFlowerMediasResponse struct {
	FlowerMedias []FlowerMedia `json:"flower_medias"`
	Count        uint64        `json:"count"`
}
