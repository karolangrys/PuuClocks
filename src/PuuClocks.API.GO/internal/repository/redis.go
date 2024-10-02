package repository

import (
	"puuclocks/internal/infrastructure"

	r "github.com/redis/go-redis/v9"
)

type Redis interface {
	infrastructure.KeyValueStore
}

type redis struct {
	infrastructure.Redis
}

type RedisConfig struct {
	Addr     string
	Password string
}

func newRedis(config RedisConfig) Redis {
	rdb := r.NewClient(&r.Options{
		Addr:     config.Addr,
		Password: config.Password,
	})

	client := infrastructure.Redis{
		Client: *rdb,
	}

	return &redis{
		client,
	}
}
