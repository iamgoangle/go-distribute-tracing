package tracing

import (
	"io"
	"log"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

// NewTracerFromEnv returns a new tracer based on the current configuration,
// using the given options, and a closer func that can be used to flush buffers before shutdown.
func NewTracerFromEnv() (opentracing.Tracer, io.Closer, error) {
	cfg, err := initFromEnv()
	if err != nil {
		log.Printf("Could not parse Jaeger env vars: %s", err.Error())
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())

		return nil, nil, err
	}

	return tracer, closer, nil
}

// InitFromEnv uses environment variables and overrides existing tracer's Configuration
//
// JAEGER_SERVICE_NAME	The service name
// JAEGER_AGENT_HOST	The hostname for communicating with agent via UDP
// JAEGER_AGENT_PORT	The port for communicating with agent via UDP
// JAEGER_ENDPOINT	The HTTP endpoint for sending spans directly to a collector, i.e. http://jaeger-collector:14268/api/traces
// JAEGER_USER	Username to send as part of "Basic" authentication to the collector endpoint
// JAEGER_PASSWORD	Password to send as part of "Basic" authentication to the collector endpoint
// JAEGER_REPORTER_LOG_SPANS	Whether the reporter should also log the spans
// JAEGER_REPORTER_MAX_QUEUE_SIZE	The reporter's maximum queue size
// JAEGER_REPORTER_FLUSH_INTERVAL	The reporter's flush interval, with units, e.g. "500ms" or "2s" (valid units)
// JAEGER_SAMPLER_TYPE	The sampler type
// JAEGER_SAMPLER_PARAM	The sampler parameter (number)
// JAEGER_SAMPLER_MANAGER_HOST_PORT	The HTTP endpoint when using the remote sampler, i.e. http://jaeger-agent:5778/sampling
// JAEGER_SAMPLER_MAX_OPERATIONS	The maximum number of operations that the sampler will keep track of
// JAEGER_SAMPLER_REFRESH_INTERVAL	How often the remotely controlled sampler will poll jaeger-agent for the appropriate sampling strategy, with units, e.g. "1m" or "30s" (valid units)
// JAEGER_TAGS	A comma separated list of name = value tracer level tags, which get added to all reported spans. The value can also refer to an environment variable using the format ${envVarName:default}, where the :default is optional, and identifies a value to be used if the environment variable cannot be found
// JAEGER_DISABLED	Whether the tracer is disabled or not. If true, the default opentracing.NoopTracer is used.
// JAEGER_RPC_METRICS	Whether to store RPC metrics
//
// see: https://github.com/jaegertracing/jaeger-client-go
func initFromEnv() (*config.Configuration, error) {
	cfg, err := config.FromEnv()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
