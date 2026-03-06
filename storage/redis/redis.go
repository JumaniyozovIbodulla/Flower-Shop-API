package redis

import (
	"context"
	"flower-shop/config"
	"flower-shop/storage"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Store struct {
	db *redis.Client
}

func New(cfg config.Config) storage.IRedisStorage {

	opt, err := redis.ParseURL(cfg.RedisUrl)

	if err != nil {
		fmt.Println("failed to connect redis")
	}

	redisClient := redis.NewClient(opt)

	return Store{
		db: redisClient,
	}
}

func (s Store) SetX(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	statusCmd := s.db.SetEx(ctx, key, value, duration)
	if statusCmd.Err() != nil {
		return statusCmd.Err()
	}
	return nil
}

func (s Store) Get(ctx context.Context, key string) interface{} {
	return s.db.Get(ctx, key).Val()
}

func (s Store) Del(ctx context.Context, key string) error {
	statusCmd := s.db.Del(ctx, key)
	if statusCmd.Err() != nil {
		return statusCmd.Err()
	}
	return nil
}
