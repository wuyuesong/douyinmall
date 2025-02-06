package clientsuite

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

type CommonGrpcClientSuite struct {
	DestServiceName    string
	DestServiceAddr    string
	CurrentServiceName string
	TracerProvider     *tracesdk.TracerProvider
}

func (s CommonGrpcClientSuite) Options() []client.Option {
	opts := []client.Option{
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithTransportProtocol(transport.GRPC),
	}

	_ = provider.NewOpenTelemetryProvider(provider.WithSdkTracerProvider(s.TracerProvider), provider.WithEnableMetrics(false))

	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		client.WithSuite(tracing.NewClientSuite()),
	)

	return opts
}
