package storage

import (
	"context"
	"flower-shop/api/models"
	"time"
)

type IStorage interface {
	CloseDB()
	Users() UserStorage
	Redis() IRedisStorage
}

type UserStorage interface {
	Create(ctx context.Context, req models.AddUser) error
	Delete(ctx context.Context, ID string) error
	Update(ctx context.Context, req models.UpdateUser) error
	UpdatePassword(ctx context.Context, req models.UpdateUserPassword) error
	GetAll(ctx context.Context, req models.GetAllUsersRequest) (models.GetAllUsersResponse, error)
}

type IRedisStorage interface {
	SetX(ctx context.Context, key string, value interface{}, duration time.Duration) error
	Get(ctx context.Context, key string) interface{}
	Del(ctx context.Context, key string) error
}
