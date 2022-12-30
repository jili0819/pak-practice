package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

var client = redis.NewClient(&redis.Options{
	Addr: ":6379",
})

func main() {
	lock01()
	//lock02()

}

// err-group
func lock01() {
	key := "key01"
	var g errgroup.Group
	ch := make(chan bool)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	_, err := client.Set(ctx, key, true, 3*time.Second).Result()
	if err != nil {
		log.Println(err.Error())
		return
	}
	// 业务处理
	g.Go(func() error {
		fmt.Println("执行业务开始")
		time.Sleep(1 * time.Second)
		fmt.Println("执行业务结束")
		ch <- true
		close(ch)
		return nil
	})
	// 锁续期
	g.Go(func() error {
		ticker := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-ticker.C:
				fmt.Println("定时续期分布式锁有效期")
				_, err := client.Expire(context.Background(), key, 3*time.Second).Result()
				if err != nil {
					log.Println(err.Error())
					return err
				}
			case <-ch:
				fmt.Println("执行业务结束,关闭续期,删除锁")
				ticker.Stop()
				_, err := client.Del(context.Background(), key).Result()
				if err != nil {
					log.Println(err.Error())
					return err
				}
				return nil
			}
		}
	})
	if err := g.Wait(); err != nil {
		fmt.Println(err)
		return
	}
}

// 使用普通方法实现
func lock02() {
	after := time.AfterFunc(time.Second*5, func() {
		fmt.Println("stop")
	})
	after.Stop()
	time.Sleep(10 * time.Second)
}
