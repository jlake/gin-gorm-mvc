package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"gin-gorm-mvc/internal/config"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client
var ctx = context.Background()

// Initialize Redis接続を初期化
func Initialize() error {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.AppConfig.Redis.Host, config.AppConfig.Redis.Port),
		Password: config.AppConfig.Redis.Password,
		DB:       config.AppConfig.Redis.DB,
	})

	// 接続テスト
	_, err := Client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}

	log.Println("Redis connection established")
	return nil
}

// Close Redis接続を閉じる
func Close() error {
	if Client != nil {
		return Client.Close()
	}
	return nil
}

// GetClient Redisクライアントを取得
func GetClient() *redis.Client {
	return Client
}

// Set キーと値を設定
func Set(key string, value interface{}, expiration time.Duration) error {
	return Client.Set(ctx, key, value, expiration).Err()
}

// Get キーから値を取得
func Get(key string) (string, error) {
	return Client.Get(ctx, key).Result()
}

// Delete キーを削除
func Delete(key string) error {
	return Client.Del(ctx, key).Err()
}

// Exists キーが存在するか確認
func Exists(key string) (bool, error) {
	result, err := Client.Exists(ctx, key).Result()
	return result > 0, err
}
