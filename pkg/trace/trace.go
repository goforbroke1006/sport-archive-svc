package trace

import (
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
	reporterhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

const endpointURI = "/api/v2/spans"

func NewTracer(zipkinHost, serviceName string, port uint16) (*zipkin.Tracer, error) {
	reporter := reporterhttp.NewReporter(zipkinHost + endpointURI)
	localEndpoint := &model.Endpoint{ServiceName: serviceName, Port: port}

	sampler, err := zipkin.NewCountingSampler(1)
	if err != nil {
		return nil, err
	}

	t, err := zipkin.NewTracer(
		reporter,
		zipkin.WithSampler(sampler),
		zipkin.WithLocalEndpoint(localEndpoint),
	)
	if err != nil {
		return nil, err
	}

	return t, err
}
