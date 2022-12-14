package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	App         string
	Environment string // dev, test, prod
	Version     string

	ServiceHost string
	HTTPPort    string

	DefaultOffset string
	DefaultLimit  string
	
	AuthorServiceGrpcHost string
	AuthorServiceGrpcPort string

	ArticleServiceGrpcHost string
	ArticleServiceGrpcPort string


}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.App = cast.ToString(getOrReturnDefaultValue("PROJECT_NAME", "Article"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "dev"))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0.0"))

	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8080"))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))

	config.AuthorServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("AUTHOR_SERVICE_GRPC_HOST", "localhost"))
	config.AuthorServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("AUTHOR_SERVICE_GRPC_PORT", ":9001"))

	config.ArticleServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("ARTICLE_SERVICE_GRPC_HOST", "localhost"))
	config.ArticleServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("ARTICLE_SERVICE_GRPC_PORT", ":9001"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
