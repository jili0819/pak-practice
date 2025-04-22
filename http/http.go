package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

func main() {
	c := gin.New()
	c.GET("/test/sse", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Stream(func(w io.Writer) bool {
			for i := 0; i < 500; i++ {
				time.Sleep(1 * time.Second)
				// 写入 SSE 数据
				_, _ = fmt.Fprintf(w, "data: %s\n\n", fmt.Sprintf("hello %d->%s", i, time.Now().Format("2006-01-02 15:04:05")))
				// 刷新流
				c.Writer.Flush()
			}
			// 写入结束
			_, _ = fmt.Fprintf(c.Writer, "data: [DONE]\n\n")
			c.Writer.Flush() // 确保输出刷新到客户端
			// 结束要返回 false
			return false
		})

	})
	c.Run(":8080")
}
