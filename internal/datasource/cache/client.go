package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/TrNgTien/new-feed-go/internal/configs"
	redis "github.com/redis/go-redis/v9"
)

var client *redis.Client

func InitRedisClient(
	cacheConfig configs.Cache,
) {

	fmt.Println("[InitRedisClient] Connecting redis!!")
	client = redis.NewClient(&redis.Options{
		Addr:     cacheConfig.Address,
		Username: cacheConfig.Username,
		Password: cacheConfig.Password,
		DB:       0, // use default DB
	})

	fmt.Println("[InitRedisClient] Connected redis!!")
}

func Get(ctx context.Context, key string) (any, error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	return val, nil
}

func Set(ctx context.Context, key string, data any, ttl time.Duration) error {

	err := client.Set(ctx, key, data, ttl).Err()

	if err != nil {
		panic(err)
	}

	return nil
}
