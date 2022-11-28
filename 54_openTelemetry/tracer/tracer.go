package tracer

//
//import "fmt"
//
//// tracerProvider is 返回一个openTelemetry TraceProvider，这里用的是jaeger
//func tracerProvider(url string) error {
//	fmt.Println("init traceProvider")
//
//	// 创建jaeger provider
//	// 可以直接连collector也可以连agent
//	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
//	if err != nil {
//		return err
//	}
//	tp = tracesdk.NewTracerProvider(
//		tracesdk.WithBatcher(exp),
//		tracesdk.WithResource(resource.NewWithAttributes(
//			semconv.SchemaURL,
//			semconv.ServiceNameKey.String(service),
//			attribute.String("environment", environment),
//			attribute.Int64("ID", id),
//		)),
//	)
//	return nil
//}
