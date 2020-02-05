package main

import (
	"context"
	"log"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/iamgoangle/go-http-tracing-kafka/pkg/tracing"
)

// TracerFunc provides a function wrapper the context to continue a span
// and also return a function to set the operation name
type TracerFunc func(funcName string) opentracing.Span

// Tag define a tag key-value type
type Tag struct {
	Key   string
	Value interface{}
}

// SpanTags defines tag array type
type SpanTags []*Tag

// Tracer provides a funcion tracing type
func Tracer(ctx context.Context, funcName string, spanTags SpanTags) TracerFunc {
	return func(n string) opentracing.Span {
		log.Println("trace: ", n)

		span, _ := opentracing.StartSpanFromContext(ctx, funcName)
		for _, t := range spanTags {
			span.SetTag(t.Key, t.Value)
		}

		// span.LogFields(
		// 	log.String("event", "HelloSarah"),
		// 	log.String("value", "Sarah Ocha"),
		// )

		return span
	}
}

func main() {
	tracer, closer := tracing.Init("demo-tracing")
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

	span := tracer.StartSpan("main program")
	span.SetTag("environtment", "production")
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)

	// 1st function call
	sp := SpanTags{
		&Tag{"channelId", 12345},
		&Tag{"campaignName", "Golf Ja"},
	}
	t := Tracer(ctx, "OnEventStart", sp)
	OnEventStart(t)

	// 2nd function call
	sp = SpanTags{
		&Tag{"channelId", 12345},
		&Tag{"campaignName", "Golf Ja"},
	}
	t = Tracer(ctx, "OnProcessor", sp)
	OnProcessor(t)

	// 3rd function call
	sp = SpanTags{
		&Tag{"channelId", 12345},
		&Tag{"campaignName", "Golf Ja"},
	}
	t = Tracer(ctx, "OnComplete", sp)
	OnComplete(t)
}

func OnEventStart(t TracerFunc) {
	span := t("logger.WithField")
	defer span.Finish()

	// do business logic...
	time.Sleep(1 * time.Second)
}

func OnProcessor(t TracerFunc) {
	span := t("logger.WithField")
	defer span.Finish()

	// do business logic...
	time.Sleep(10 * time.Second)
}

func OnComplete(t TracerFunc) {
	span := t("logger.WithField")
	defer span.Finish()

	// do business logic...
	time.Sleep(10 * time.Second)
}
