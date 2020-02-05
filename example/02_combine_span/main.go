package main

import (
	"context"
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/iamgoangle/go-http-tracing-kafka/pkg/tracing"
)

func main() {
	tracer, closer := tracing.Init("tracer-context")
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

	span := tracer.StartSpan("main program")
	span.SetTag("channelID", 112)
	span.SetTag("campaignName", "Golf Ja")
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)

	OnEventStart(ctx, "John")
	OnProcessor(ctx)
	OnComplete(ctx)
}

func OnEventStart(ctx context.Context, name string) {
	fmt.Println("call OnEventStart")
	span, _ := opentracing.StartSpanFromContext(ctx, "OnEventStart")
	defer span.Finish()

	time.Sleep(1 * time.Second)
}

func OnProcessor(ctx context.Context) {
	fmt.Println("call OnProcessor")
	span, _ := opentracing.StartSpanFromContext(ctx, "OnProcessor")
	defer span.Finish()

	time.Sleep(10 * time.Second)
}

func OnComplete(ctx context.Context) {
	fmt.Println("call OnComplete")

	span, _ := opentracing.StartSpanFromContext(ctx, "OnComplete")
	defer span.Finish()

	time.Sleep(10 * time.Second)
}
