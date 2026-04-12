package storage

import (
	"context"
	"flower-shop/api/models"
	"time"
)

type IStorage interface {
	CloseDB()
	Users() UserStorage
	Roles() RoleStorage
	Permissions() PermissionStorage
	RolePermissions() RolePermissionsStorage
	UserRoles() UserRolesStorage
	Redis() IRedisStorage
}

type UserStorage interface {
	Create(ctx context.Context, req models.AddUser) error
	Delete(ctx context.Context, ID string) error
	Update(ctx context.Context, req models.UpdateUser) error
	UpdatePassword(ctx context.Context, req models.UpdateUserPassword) error
	GetAll(ctx context.Context, req models.GetAllUsersRequest) (models.GetAllUsersResponse, error)
}

type RoleStorage interface {
	Create(ctx context.Context, req models.AddRole) error
	Delete(ctx context.Context, ID string) error
	Update(ctx context.Context, req models.UpdateRole) error
	GetAll(ctx context.Context, req models.GetAllRolesRequest) (models.GetAllRolesResponse, error)
}

type PermissionStorage interface {
	Create(ctx context.Context, req models.AddPermission) error
	Delete(ctx context.Context, ID string) error
	Update(ctx context.Context, req models.UpdatePermission) error
	GetAll(ctx context.Context, req models.GetAllPermissionsRequest) (models.GetAllPermissionsResponse, error)
}

type RolePermissionsStorage interface {
	Create(ctx context.Context, req models.RolePermission) error
	Delete(ctx context.Context, req models.RolePermission) error
}

type UserRolesStorage interface {
	Create(ctx context.Context, req models.UserRole) error
	Delete(ctx context.Context, req models.UserRole) error
	Update(ctx context.Context, req models.UserRole) error
}

type IRedisStorage interface {
	SetX(ctx context.Context, key string, value interface{}, duration time.Duration) error
	Get(ctx context.Context, key string) interface{}
	Del(ctx context.Context, key string) error
}
