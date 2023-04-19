package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func TimeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		fmt.Println("start time: ", startTime.Format(time.DateTime))
		c.Next()
		fmt.Println("end time: ", time.Since(startTime))
		return
	}
}
