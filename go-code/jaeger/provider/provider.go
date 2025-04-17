package provider

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/log"
	"io"
)

func NewProvider(service string) (io.Closer, error) {
	setting := &config.Configuration{
		ServiceName: service,
		Sampler:     &config.SamplerConfig{},
		Reporter:    &config.ReporterConfig{},
	}

	if tracer, closer, err := setting.NewTracer(config.Logger(log.StdLogger)); err != nil {
		return nil, err
	} else {
		opentracing.SetGlobalTracer(tracer)
		return closer, nil
	}
}
