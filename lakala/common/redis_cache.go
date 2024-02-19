package common

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisCache struct {
	ICache
	cache redis.UniversalClient
}

func NewRedisCache(cache redis.UniversalClient) *RedisCache {
	return &RedisCache{
		cache: cache,
	}
}

func (r *RedisCache) GetAccessToken(ctx context.Context, key string) (token string, err error) {
	if token, err = r.cache.Get(ctx, key).Result(); err != nil && err != redis.Nil {
		return "", err
	}
	err = nil
	return
}

func (r *RedisCache) SetAccessToken(ctx context.Context, key, token string, expireTime time.Duration) (err error) {
	return r.cache.Set(ctx, key, token, expireTime).Err()
}

func (r *RedisCache) DelAccessToken(ctx context.Context, key string) (err error) {
	return r.cache.Del(ctx, key).Err()
}
