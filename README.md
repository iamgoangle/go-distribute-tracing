# Workshop :free:

- Create `docker-compose` for setting up the project.
- Learning about Jaeger Architecture <https://www.jaegertracing.io/docs/1.11/architecture/>
- Learning about `span` `trace` and familiar with `baggage`.
- Learning about [Environment Variable in Go Jaeger Client](https://github.com/jaegertracing/jaeger-client-go#environment-variables)
- Create `logger` for logging with tracing features.
- Create go for tracing for `all in-bound HTTP request`.
- Create go for `tracing for Kafka producer`.
- Create go for `tracing for in-bound Kafka consumer`.

### Getting started

```sh
docker-compose up --build
```

### Jaeger UI

http://localhost:16686

## Storage

With this demo All-in-one `docker-compose` use an in memory storage component.

## Prometheus

All Jaeger backend components expose metrics by default

<https://www.jaegertracing.io/docs/1.11/features/#observability>

### Jaeger Address

- Distributed transaction monitoring.
- Performance and latency optimization.
- Root cause analysis.
- Service dependency analysis.
- Distributed context propagation.
