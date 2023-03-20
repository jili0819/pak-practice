package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/jili/pkg-practice/lock/sync_map"
	"log"
	"sync"
	"time"
)

const DefaultLockTime = time.Second * 3

var (
	rdb  *redis.Client
	once sync.Once
)

func GetRedisClient() *redis.Client {
	if rdb != nil {
		return rdb
	}
	once.Do(func() {
		rdb = newRedisClient()
	})
	return rdb
}

func newRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: ":6379",
	})
}

func Lock(ctx context.Context, mx *sync_map.KeyedMutex, key string) error {
	unlock := mx.Lock(key)
	defer unlock()
	_, err := GetRedisClient().SetNX(ctx, key, 1, DefaultLockTime).Result()
	if err != nil {
		log.Println(err.Error())
	}
	return err
}

func UnLock(ctx context.Context, key string) error {
	_, err := GetRedisClient().Del(ctx, key).Result()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func LockExpire(ctx context.Context, key string) error {
	_, err := GetRedisClient().Expire(ctx, key, time.Second*1).Result()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return err
}
