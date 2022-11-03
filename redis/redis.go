package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	client := redis.NewClient(
		&redis.Options{
			Addr: ":6379",
		})
	//client.SetNX(ctx, "aa", "fwfwfwefwe", 1*time.Hour)
	err := client.Get(ctx, "aaa").Err()
	fmt.Println(err)
	//fmt.Println(err)
}
