package main

import (
	"context"
	"flower-shop/api"
	"flower-shop/config"
	"flower-shop/pkg/logger"
	"flower-shop/service"
	"flower-shop/storage/postgres"
	"flower-shop/storage/redis"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func envPort(host string, port string) string {
	return host + ":" + port
}

func main() {
	cfg := config.Load()
	redis := redis.New(cfg)
	log := logger.New(cfg.ServiceName)

	store, err := postgres.New(context.Background(), cfg, redis)
	if err != nil {
		log.Error("error while connecting db, err: ", logger.Error(err))
		return
	}

	defer store.CloseDB()

	service := service.New(store, log)

	c := api.New(store, service, log)

	if cfg.GinMode == "" {
		cfg.GinMode = os.Getenv("GIN_MODE")
	}

	if cfg.GinMode == "debug" || cfg.GinMode == "release" || cfg.GinMode == "test" {
		gin.SetMode(cfg.GinMode)

	} else {
		fmt.Printf("Invalid GIN_MODE: %s (valid modes: debug, release, test)", cfg.GinMode)
	}

	address := envPort(cfg.ProdHost, cfg.ProdPort)
	c.Run(address)
}
