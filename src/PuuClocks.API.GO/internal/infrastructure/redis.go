package infrastructure

import (
	"context"
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type Redis struct {
	Client redis.Client
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	res, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("redis - couldn't get value for %s key: %w", key, err)
	}
	return res, nil
}

func (r *Redis) Set(ctx context.Context, key, value string, exp time.Duration) error {
	_, err := r.Client.Set(ctx, key, value, exp).Result()
	if err != nil {
		return fmt.Errorf("redis - couldn't set value for %s key: %w", key, err)
	}
	return nil
}

func (r *Redis) Health(ctx context.Context) error {
	_, err := r.Client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("redis - is not healthy: %w", err)
	}
	return nil
}
