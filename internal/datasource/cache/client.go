package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/TrNgTien/new-feed-go/internal/configs"
	redis "github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(
	cacheConfig configs.Cache,
) *RedisClient {

	fmt.Println("[InitRedisClient] Connecting redis!!")
	client := redis.NewClient(&redis.Options{
		Addr:     cacheConfig.Address,
		Username: cacheConfig.Username,
		Password: cacheConfig.Password,
		DB:       0, // use default DB
	})

	fmt.Println("[InitRedisClient] Connected redis!!")
	return &RedisClient{client: client}
}

func (c *RedisClient) Get(ctx context.Context, key string) (any, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	return val, nil
}

func (c *RedisClient) Set(ctx context.Context, key string, data any, ttl time.Duration) error {

	err := c.client.Set(ctx, key, data, ttl).Err()

	if err != nil {
		panic(err)
	}

	return nil
}
