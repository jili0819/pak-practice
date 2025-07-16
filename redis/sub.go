package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jili/pkg-practice/gin/middleware"
	"io"
	"sync"
	"time"
)

var (
	redisClient *redis.Client
)

func init() {
	redisClient = redis.NewClient(
		&redis.Options{
			Addr: ":6379",
		})
}

type AA struct {
	Com string `json:"com"`
	B   string `json:"b"`
}

func main() {
	ss := []AA{
		{
			Com: "com",
			B:   "b",
		},
		{
			Com: "com",
			B:   "c",
		},
	}
	ssss := make(map[string][]AA)
	for _, s := range ss {
		if _, ok := ssss[s.Com]; !ok {
			ssss[s.Com] = []AA{}
		}
		ssss[s.Com] = append(ssss[s.Com], s)
	}
	fmt.Println(ssss)

	return

	g := gin.Default()
	g.Use(middleware.TimeMiddleware())
	g.GET("/", func(c *gin.Context) {
		return
	})
	g.GET("/sse", func(c *gin.Context) {
		ctx, newCancel := context.WithCancel(context.Background())
		defer newCancel()
		redisClient.Set(ctx, "sse", "running", 24*time.Hour)
		defer redisClient.Set(ctx, "sse", "stop", 24*time.Hour)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			i := 0
			for {
				if i > 60 {
					break
				}
				if err := redisClient.Publish(ctx, "sse", fmt.Sprintf("data: [heartbeat] %d\n\n", i)).Err(); err != nil {
					fmt.Println("publish err", err)
				}
				time.Sleep(1 * time.Second)
				i++
			}
		}()
		wg.Wait()
	})
	g.GET("/get", func(c *gin.Context) {
		// 设置 SSE Header
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Stream(func(w io.Writer) bool {
			sub := redisClient.Subscribe(context.Background(), "sse")
			defer func() {
				// 取消订阅
				_ = sub.Unsubscribe(context.Background(), "sse")
			}()
			fmt.Println("------------")
			ticker := time.NewTicker(5 * time.Second)
			defer ticker.Stop()
			for {
				// 5秒一次校验redis中指定key的value
				select {
				case <-ticker.C:
					// 发送数据
					_, _ = fmt.Fprintf(w, fmt.Sprintf("data: [heartbeat]\n\n"))
					c.Writer.Flush()
				case msg := <-sub.Channel():
					// 发送数据
					_, _ = fmt.Fprintf(w, fmt.Sprintf("data: %s\n\n", msg.Payload))
					c.Writer.Flush()
				default:
					// 这里是为了避免死循环
				}
			}
			return false
		})
	})
	g.Run(":8080")
}
