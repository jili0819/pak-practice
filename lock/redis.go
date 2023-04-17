package main

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	LockPrefix        = "{redis_lock}_"
	DefaultExpiration = 5
)

type RedisLock struct {
	ctx        context.Context
	key        string
	value      string
	client     *redis.Client
	expiration time.Duration
	cancelFunc context.CancelFunc
}

func NewRedisLock(ctx context.Context, client *redis.Client, key string, value string) *RedisLock {
	instancLock := &RedisLock{
		ctx:        ctx,
		key:        LockPrefix + key,
		value:      value,
		client:     client,
		expiration: time.Duration(DefaultExpiration) * time.Second,
	}
	return instancLock
}

func (c *RedisLock) SetExpiration(expiration time.Duration) *RedisLock {
	c.expiration = expiration
	return c
}

func (c *RedisLock) TryLock() (success bool, err error) {
	success, err = c.client.SetNX(c.ctx, c.key, c.value, c.expiration).Result()
	if err != nil {
		return
	}
	ctx, cancel := context.WithCancel(c.ctx)
	c.cancelFunc = cancel
	c.renew(ctx)
	return
}

func (c *RedisLock) Lock() error {
	for {
		success, err := c.TryLock()
		if err != nil {
			return err
		}
		if success {
			return nil
		}
	}
	return nil
}

func (c *RedisLock) UnLock() (err error) {
	c.cancelFunc() //cancel renew goroutine
	return
}

func (c *RedisLock) renew(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(c.expiration / 3)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				c.client.Expire(ctx, c.key, c.expiration).Result()
			}
		}
	}()
}
