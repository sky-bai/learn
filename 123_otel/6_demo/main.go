package main

import (
	"context"
	"go.opentelemetry.io/otel"
)

var traceName = "gin"

func main() {
	// 注意！这里仅是测试的一个函数 并不包括完全的过程，仅用于测试部署，不涉及链路追踪知识
	// 父span
	newCtx, span := otel.Tracer(traceName).Start(ctx, "getUserID", trace.WithSpanKind(trace.SpanKindServer))
	defer span.End()

	child1(newCtx)

}

func child1(ctx context.Context) {
	//子span一
	_, span := otel.Tracer(traceName).Start(ctx, "getUserInfo", trace.WithSpanKind(trace.SpanKindServer))
	span.End()
}

func child2(ctx context.Context) {
	//子span二
	_, span := otel.Tracer(traceName).Start(ctx, "getMember", trace.WithSpanKind(trace.SpanKindServer))
	span.End()
}
