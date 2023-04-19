package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"net/http"
)

var (
	SpanCTX = "span-ctx"
)

func Jaeger() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.GetHeader("uber-trace-id")
		var span opentracing.Span
		if traceId != "" {
			var err error
			span, err = GetParentSpan(c.FullPath(), traceId, c.Request.Header)
			if err != nil {
				return
			}
		} else {
			span = StartSpan(opentracing.GlobalTracer(), c.FullPath())
		}
		defer span.Finish()
		c.Set(SpanCTX, opentracing.ContextWithSpan(c, span))
		c.Next()
	}
}

//
//func main() {
//	g := gin.Default()
//	g.Use(Jaeger())
//	g.GET("/", func(c *gin.Context) {
//		spanCtxInterface, _ := c.Get(SpanCTX)
//		var spanCtx context.Context
//		spanCtx = spanCtxInterface.(context.Context)
//		//创建子span
//		span, _ := WithSpan(spanCtx, "JaegerGet")
//		defer span.Finish() //结束后调用完成
//	})
//	g.POST("/post", func(c *gin.Context) {
//		spanCtxInterface, _ := c.Get(SpanCTX)
//		var spanCtx context.Context
//		spanCtx = spanCtxInterface.(context.Context)
//		//创建子span
//		span, _ := WithSpan(spanCtx, "JaegerPost")
//		defer span.Finish() //结束后调用完成
//		carrier, _ := GetCarrier(span)
//		client := &http.Client{}
//		req, _ := http.NewRequest("GET", "http://127.0.0.1:8080/jaegerTest", bytes.NewReader([]byte{}))
//		req.Header.Add("User-Agent", "myClient")
//		_ = carrier.ForeachKey(func(key, val string) (err error) {
//			req.Header.Add(key, val)
//			return
//		})
//		resp, _ := client.Do(req)
//		defer resp.Body.Close()
//	})
//}

func GetDefaultConfig() *config.Configuration {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
	}
	return cfg
}
func init() {
	jaegerConfig := GetDefaultConfig()
	InitJaeger("go-framework-demo", jaegerConfig)
}

/*
*
初始化
*/
func InitJaeger(service string, cfg *config.Configuration) (opentracing.Tracer, io.Closer) {
	cfg.ServiceName = service
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("Error: connot init Jaeger: %v\\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}
func StartSpan(tracer opentracing.Tracer, name string) opentracing.Span {
	//设置顶级span
	span := tracer.StartSpan(name)
	return span
}
func WithSpan(ctx context.Context, name string) (opentracing.Span, context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, name)
	return span, ctx
}
func GetCarrier(span opentracing.Span) (opentracing.HTTPHeadersCarrier, error) {
	carrier := opentracing.HTTPHeadersCarrier{}
	err := span.Tracer().Inject(span.Context(), opentracing.HTTPHeaders, carrier)
	if err != nil {
		return nil, err
	}
	return carrier, nil
}
func GetParentSpan(spanName string, traceId string, header http.Header) (opentracing.Span, error) {
	carrier := opentracing.HTTPHeadersCarrier{}
	carrier.Set("uber-trace-id", traceId)
	tracer := opentracing.GlobalTracer()
	wireContext, err := tracer.Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(header),
	)
	parentSpan := opentracing.StartSpan(
		spanName,
		ext.RPCServerOption(wireContext))
	if err != nil {
		return nil, err
	}
	return parentSpan, err
}
