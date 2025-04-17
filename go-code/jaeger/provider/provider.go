package provider

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/log"
	"io"
)

func NewProvider(service string) (io.Closer, error) {
	setting := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst, // 모든 Span을 샘플링(Span을 처리한다) 한다.
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
			//LocalAgentHostPort: "localhost:6820",
		},
	}
	if tracer, closer, err := setting.NewTracer(config.Logger(log.StdLogger)); err != nil {
		return nil, err
	} else {
		opentracing.SetGlobalTracer(tracer)
		return closer, nil
	}
}
