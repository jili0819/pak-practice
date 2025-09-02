package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

type TypeName struct {
	Bs []string `json:"bs"`
}
type A struct {
	Name     string   `json:"name" xml:"name"`
	Password string   `json:"password" xml:"password"`
	T        TypeName `json:"t" xml:"t"`
}

type Req struct {
	Name string `form:"name" json:"name"`
}

func main() {
	g := gin.Default()
	g.POST("/", func(c *gin.Context) {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		c.Abort()
		return
	})
	g.POST("/notify", func(c *gin.Context) {
		fmt.Println(c.Request.URL)
		fmt.Println("-------")
		fmt.Println(c.Request.RequestURI)
		fmt.Println("-------")
		fmt.Println(c.Request.Header)
		fmt.Println("-------")
		cc, _ := c.GetRawData()
		fmt.Println(string(cc))
		return
	})
	g.GET("/sse", func(c *gin.Context) {
		_, newCancel := context.WithCancel(context.Background())
		// 设置 SSE Header
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		i := 0
		c.Stream(func(w io.Writer) bool {
			for {
				if i > 100 {
					newCancel()
					break
				}
				fmt.Println("i:", i)
				time.Sleep(1 * time.Second)
				// 发送数据
				_, _ = fmt.Fprintf(w, "data: [heartbeat]\n\n")
				c.Writer.Flush()
				i++
			}
			return false
		})
	})
	g.GET("/get", func(c *gin.Context) {
		data := "Hello, go-stress-testing! \n"
		bb := make([]int, 0, 1000000)
		fmt.Println(bb)
		c.Writer.Header().Set("Server", "golang")
		c.Writer.Write([]byte(data))
		return
	})
	g.Run(":4000")
}

/*import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

const (
	service     = "trace-demo"
	environment = "production"
	id          = 1
)

// tracerProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(service),
			attribute.String("environment", environment),
			attribute.Int64("ID", id),
		)),
	)
	return tp, nil
}

func main() {
	tp, err := tracerProvider("http://localhost:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Cleanly shutdown and flush telemetry when the application exits.
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}(ctx)

	tr := tp.Tracer("component-main")

	ctx, span := tr.Start(ctx, "foo")
	defer span.End()

	bar(ctx)
}

func bar(ctx context.Context) {
	// Use the global TracerProvider.
	tr := otel.Tracer("component-bar")
	_, span := tr.Start(ctx, "bar")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()

	// Do bar...
}
*/
