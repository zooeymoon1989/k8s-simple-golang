package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	"k8s-simple-golang/actions"
	"k8s-simple-golang/config"
	"net/http"
	"os"
)

func traceProvider(url string) (*tracesdk.TracerProvider, error) {
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
			semconv.ServiceName("k8s-simple-golang"),
			attribute.String("environment", "production"),
		)),
	)
	return tp, nil
}

func main() {
	v := config.Config{}
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	//tp, err := traceProvider(v.JaegerAddr)
	//if err != nil {
	//	panic(err)
	//}

	//// Register our TracerProvider as the global so any imported
	//// instrumentation in the future will default to using it.
	//otel.SetTracerProvider(tp)
	//
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	//// Cleanly shutdown and flush telemetry when the application exits.
	//defer func(ctx context.Context) {
	//	// Do not make the application hang when it is shutdown.
	//	ctx, cancel = context.WithTimeout(ctx, time.Second*5)
	//	defer cancel()
	//	if err := tp.Shutdown(ctx); err != nil {
	//		log.Fatal(err)
	//	}
	//}(ctx)

	//tr := tp.Tracer("component-main")
	//
	//ctx, span := tr.Start(ctx, "foo")
	//defer span.End()

	r := setupRouter()
	err = viper.Unmarshal(&v)
	if err != nil {
		panic(fmt.Errorf("fatal error Unmarshal file: %s", err))
	}
	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		id, _ := os.Hostname()
		host := c.Request.Host
		c.String(http.StatusOK, fmt.Sprintf("hello from go-gin , the host address is: %s, and the hostname is: %s", host, id))
	})
	r.GET("/reviews/:id", actions.GetReviews)
	r.GET("/reviews", actions.GetReviews)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}
