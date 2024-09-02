package cache

import (
	"context"
	"time"

	"github.com/TrNgTien/new-feed-go/internal/configs"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(cacheConfig configs.Cache, l *zap.Logger) (*RedisClient, func(), error) {
	l.Info("[InitRedisClient] Connecting to Redis")

	client := redis.NewClient(&redis.Options{
		Addr:     cacheConfig.Address,
		Username: cacheConfig.Username,
		Password: cacheConfig.Password,
		DB:       0, // use default DB
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		l.Error("[InitRedisClient] Error connecting to Redis")
		return nil, nil, err
	}

	l.Info("[InitRedisClient] Connected to Redis")

	cleanup := func() {
		if err := client.Close(); err != nil {
			l.Error("[InitRedisClient] Error closing Redis connection")
		}
	}

	return &RedisClient{client: client}, cleanup, nil
}

func (c *RedisClient) Get(ctx context.Context, key string) (any, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (c *RedisClient) Set(ctx context.Context, key string, data any, ttl time.Duration) error {
	return c.client.Set(ctx, key, data, ttl).Err()
}
