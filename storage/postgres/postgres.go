package postgres

import (
	"context"
	"flower-shop/config"
	"flower-shop/storage"
	"flower-shop/storage/redis"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	Pool  *pgxpool.Pool
	cfg   config.Config
	redis storage.IRedisStorage
}

func New(ctx context.Context, cfg config.Config, redis storage.IRedisStorage) (storage.IStorage, error) {

	pgxPoolConfig, err := pgxpool.ParseConfig(cfg.DatabaseUrl)

	if err != nil {
		return nil, err
	}

	pgxPoolConfig.MaxConns = 50
	pgxPoolConfig.MaxConnLifetime = time.Hour

	newPool, err := pgxpool.NewWithConfig(ctx, pgxPoolConfig)
	if err != nil {
		return nil, err
	}

	return Store{
		Pool:  newPool,
		redis: redis,
		cfg:   cfg,
	}, nil
}
func (s Store) CloseDB() {
	s.Pool.Close()
}

func (s Store) Users() storage.UserStorage {
	newUser := NewUser(s.Pool)
	return &newUser
}

func (s Store) Roles() storage.RoleStorage {
	newRole := NewRole(s.Pool)
	return &newRole
}

func (s Store) Permissions() storage.PermissionStorage {
	newPermission := NewPermission(s.Pool)
	return &newPermission
}

func (s Store) RolePermissions() storage.RolePermissionsStorage {
	newRolePermissions := NewRolePermissions(s.Pool)
	return &newRolePermissions
}


func (s Store) UserRoles() storage.UserRolesStorage {
	newUserRoles := NewUserRoles(s.Pool)
	return &newUserRoles
}

func (s Store) Redis() storage.IRedisStorage {
	return redis.New(s.cfg)
}
