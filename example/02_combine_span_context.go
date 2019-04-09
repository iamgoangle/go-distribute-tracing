package main

import (
	"context"

	"github.com/iamgoangle/go-http-tracing-kafka/internal/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

func main() {
	tracer, closer := tracing.Init("my-hello-world")
	defer closer.Close()

	// GlobalTracer returns the global singleton `Tracer` implementation
	opentracing.SetGlobalTracer(tracer)

	// Create span
	span := tracer.StartSpan("say-my-name")
	span.SetTag("value", "ShaZam")
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	HelloJohn(ctx)
	HelloSarah(ctx)
}

func HelloJohn(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "HelloJohn")
	defer span.Finish()

	span.LogFields(
		log.String("event", "HelloJohn"),
		log.String("value", "John Ocha"),
	)
}

func HelloSarah(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "HelloSarah")
	defer span.Finish()

	span.LogFields(
		log.String("event", "HelloSarah"),
		log.String("value", "Sarah Ocha"),
	)
}
