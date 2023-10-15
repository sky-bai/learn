package main

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"io"
	"log"
	"os"

	"os/signal"
)

//回溯一点，跟踪是一种遥测，表示服务正在完成的工作。

// 要在跟踪器中唯一标识你的应用程序，你需要在app.go中创建一个包名常量。
// name is the Tracer name used to identify this instrumentation library.
const name = "fib"

func main() {

	l := log.New(os.Stdout, "", 0)

	// Write telemetry data to a file.
	f, err := os.Create("./123_otel/traces.txt")
	if err != nil {
		l.Fatal(err)
	}
	defer f.Close()

	//exp, err := newExporter(f)
	//if err != nil {
	//	l.Fatal(err)
	//}
	//
	//tp := trace.NewTracerProvider(
	//	// 设置需要投入的exporter WithBatcher registers the exporter
	//	trace.WithBatcher(exp),
	//	trace.WithResource(newResource()),
	//)
	tp, err := initTracer("http://47.107.47.161:14268/api/traces")
	if err != nil {
		l.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			l.Fatal(err)
		}
	}()
	otel.SetTracerProvider(tp)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error)
	app := NewApp(os.Stdin, l)
	go func() {
		errCh <- app.Run(context.Background())
	}()

	select {
	case <-sigCh:
		l.Println("\ngoodbye")
		return
	case err := <-errCh:
		if err != nil {
			l.Fatal(err)
		}
	}
}

// Fibonacci returns the n-th Fibonacci number.
func Fibonacci(n uint) (uint64, error) {
	if n <= 1 {
		return uint64(n), nil
	}
	if n > 93 {
		return 0, fmt.Errorf("unsupported Fibonacci number %d: too large", n)
	}
	var n2, n1 uint64 = 0, 1
	for i := uint(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1, nil
}

// newExporter returns a console exporter.
func newExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
}

// 设置全局trace
func initTracer(url string) (*trace.TracerProvider, error) {
	// 创建 Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := trace.NewTracerProvider(
		// 将基于父span的采样率设置为100%
		trace.WithSampler(trace.ParentBased(trace.TraceIDRatioBased(1.0))),
		// 始终确保在生产中批量处理
		trace.WithBatcher(exp),
		// 在资源中记录有关此应用程序的信息
		trace.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String("kratos-trace"),
			attribute.String("exporter", "jaeger"),
			attribute.Float64("float", 312.23),
		)),
	)
	otel.SetTracerProvider(tp)
	return tp, nil
}

// OpenTelemetry uses a Resource to represent the entity producing telemetry

// newResource returns a resource describing this application.
func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("fib"),
			semconv.ServiceVersion("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)
	return r
}

// 这就是使用TracerProvider的地方。它是一个中心点，仪表将从这些示踪器获取遥测数据，并将这些数据导入输出管道。
