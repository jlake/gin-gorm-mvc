package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	Redis    RedisConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type ServerConfig struct {
	APIPort   string
	FrontPort string
	AdminPort string
}

var AppConfig *Config

func Load() error {
	// .envファイルを読み込む（存在しない場合はスキップ）
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	AppConfig = &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", ""),
			Database: getEnv("DB_NAME", "gin_gorm_mvc"),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       0,
		},
		Server: ServerConfig{
			APIPort:   getEnv("API_PORT", "8080"),
			FrontPort: getEnv("FRONT_PORT", "8081"),
			AdminPort: getEnv("ADMIN_PORT", "8082"),
		},
	}

	return nil
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Database,
	)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
