package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	PostServiceHost   string
	PostServicePort   string
	UserServiceHost   string
	UserServicePort   string
	CommentServiceHost string
	CommentServicePort string

	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	LogLevel         string
}

func Load() Config {
	c := Config{}
	
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToString(getOrReturnDefault("POSTGRES_PORT", "5432"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "ismoiljon12"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "12"))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "comment_db"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))

	c.CommentServiceHost = cast.ToString(getOrReturnDefault("COMMIT_SERVICE_HOST","localhost"))
	c.CommentServicePort = cast.ToString(getOrReturnDefault("COMMIT_SERVICE_PORT",":8001"))

	c.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "localhost"))
	c.PostServicePort = cast.ToString(getOrReturnDefault("POST_SERVICE_PORT", "9090"))

	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "localhost"))
	c.UserServicePort = cast.ToString(getOrReturnDefault("USER_SERVICE_PORT", "8081"))
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
