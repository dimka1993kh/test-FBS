package repository

import (
	"context"
	"net"

	"github.com/caarlos0/env/v6"
	"github.com/go-redis/redis/v8"
)

//go:generate mockgen -package=mocks -source=./../../internal/repository/redis.go -destination=./../../mocks/Repository.go
type RedisInterface interface {
	HSet(ctx context.Context, key string, value interface{}) error
	HGet(ctx context.Context, key string) (string, error)
	ClearAllCache(ctx context.Context)
}

type Config struct {
	Host     string `env:"REDIS_HOST" envDefault:"0.0.0.0"`
	Port     string `env:"REDIS_PORT" envDefault:"1000"`
	Password string `env:"REDIS_PASSWORD" envDefault:""`
	DB       int    `env:"DEFAULT_DB" envDefault:"0"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

type Redis struct {
	client *redis.Client
}

func (r *Redis) HGet(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *Redis) HSet(ctx context.Context, key string, value interface{}) error {
	return r.client.Set(ctx, key, value, 0).Err()
}

func (r *Redis) ClearAllCache(ctx context.Context) {
	r.client.FlushDB(ctx)
}

func NewRedis(cfg *Config) *Redis {
	return &Redis{
		redis.NewClient(&redis.Options{
			Addr:     net.JoinHostPort(cfg.Host, cfg.Port),
			Password: cfg.Password,
			DB:       cfg.DB,
		}),
	}
}
