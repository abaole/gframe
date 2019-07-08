package gframe

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/zipkin"
)

//var Tracer Tracer
type tracerConfig struct {
	AppName string `json:"app_name"`
	Address string `json:"addr"`
}

var Tracer *opentracing.Tracer

func InitTracing(opt tracerConfig) error {
	propagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
	sender, _ := jaeger.NewUDPTransport(opt.Address, 0)
	tracer, _ := jaeger.NewTracer(
		opt.AppName,
		jaeger.NewConstSampler(true),
		jaeger.NewRemoteReporter(sender),
		jaeger.TracerOptions.Injector(opentracing.HTTPHeaders, propagator),
		jaeger.TracerOptions.Extractor(opentracing.HTTPHeaders, propagator),
		jaeger.TracerOptions.ZipkinSharedRPCSpan(true),
	)
	opentracing.SetGlobalTracer(tracer)
	Tracer = &tracer

	return nil
}

func closeTracer() error {

	return nil
}
