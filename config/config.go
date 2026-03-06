package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	ServiceName string
	DatabaseUrl string
	RedisUrl    string
	GinMode     string
	ProdPort    string
	ProdHost    string
}

func Load() Config {

	if err := godotenv.Load(); err != nil {
		fmt.Println("no .env file: ", err)
	}

	cfg := Config{}

	cfg.DatabaseUrl = cast.ToString(getOrReturnDefault("DATABASE_URL", ""))
	cfg.RedisUrl = cast.ToString(getOrReturnDefault("REDIS_URL", ""))
	cfg.ServiceName = cast.ToString(getOrReturnDefault("SERVICE_NAME", ""))
	cfg.GinMode = cast.ToString(getOrReturnDefault("GIN_MODE", ""))
	cfg.ProdPort = cast.ToString(getOrReturnDefault("PORT", "3000"))
	cfg.ProdHost = cast.ToString(getOrReturnDefault("HOST", ""))

	return cfg
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {

	if os.Getenv(key) == "" {
		return defaultValue
	}
	return os.Getenv(key)
}
