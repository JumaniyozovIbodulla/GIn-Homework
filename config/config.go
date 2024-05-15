package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost     string
	PostgresPort     int
	PostgresPassword string
	PostgresUser     string
	PostgresDatabase string
	ServiceName      string
	RedisHost        string
	RedisPort        string
	RedisPassword    string
}

func Load() Config {

	if err := godotenv.Load("/home/acer/golang/lms/.env"); err != nil {
		fmt.Println("no .env file: ", err)
	}
	cfg := Config{}

	cfg.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	cfg.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	cfg.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "gin1"))
	cfg.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "golang"))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "golang"))
	cfg.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", "localhost"))
	cfg.RedisPort = cast.ToString(getOrReturnDefault("REDIS_PORT", "6379"))
	cfg.RedisPassword = cast.ToString(getOrReturnDefault("REDIS_PASSWORD", "password"))
	cfg.ServiceName = "LMS"

	return cfg
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {

	if os.Getenv(key) == "" {
		return defaultValue
	}
	return os.Getenv(key)
}
