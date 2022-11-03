package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//provider := trace.NewNoopTracerProvider()
	//r.Use(middleware.GinMiddleware("test_tracer", tracer.WithTracerProvider(provider)))
	r.POST("/ctx", func(c *gin.Context) {
		//ctx := c.Request.Context()
		// 通过http header，提取span元数据信息
		/*span := trace.SpanFromContext(otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(c.Request.Header)))
		defer span.End()
		fmt.Println(span.SpanContext().TraceID(), span.SpanContext().SpanID())*/
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = r.Run(":9090")
}
