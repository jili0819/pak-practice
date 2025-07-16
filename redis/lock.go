package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	uuid2 "github.com/google/uuid"
	"time"
)

const unlockScript = `
if redis.call("get", KEYS[1]) == ARGV[1] then
    return redis.call("del", KEYS[1])
else
    return 0
end`

func main() {
	ctx := context.Background()
	client := redis.NewClient(
		&redis.Options{
			Addr: ":6379",
		})
	lockKey := "lock_key"
	uuid := uuid2.New().String()
	xx, err := client.SetNX(ctx, lockKey, uuid, 3*time.Second).Result()
	fmt.Println("lock", xx, err)

	xxi, err := client.Eval(ctx, unlockScript, []string{lockKey}, uuid).Result()
	fmt.Println("unlock", xxi, err)
}
