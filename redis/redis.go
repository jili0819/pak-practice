package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var a map[int]int

func main() {
	ctx := context.Background()
	client := redis.NewClient(
		&redis.Options{
			Addr: ":6379",
		})
	xx, err := client.SetNX(ctx, "bb", "fwfwfwefwe", 3*time.Second).Result()
	fmt.Println(xx, err)

	//err := client.HIncrBy(context.Background(), "test0101", strconv.FormatUint(uint64(101), 10), 1).Err()
	//fmt.Println(err)
	//fmt.Println(client.HGet(context.Background(), "test0101", strconv.FormatUint(uint64(101), 10)).Int())

}
