package operator

import (
	"context"

	"google.golang.org/grpc/health/grpc_health_v1"
)

type healthCheckImpl struct{}

// Check implements grpc_health_v1.HealthServer.
func (h *healthCheckImpl) Check(context.Context, *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	panic("unimplemented")
}

// Watch implements grpc_health_v1.HealthServer.
func (h *healthCheckImpl) Watch(*grpc_health_v1.HealthCheckRequest, grpc_health_v1.Health_WatchServer) error {
	panic("unimplemented")
}

var _ grpc_health_v1.HealthServer = &healthCheckImpl{}
