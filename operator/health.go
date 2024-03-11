package operator

import (
	"context"

	"google.golang.org/grpc/health/grpc_health_v1"
)

type healthCheckImpl struct{}

// Check implements grpc_health_v1.HealthServer.
func (h *healthCheckImpl) Check(context.Context, *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

// Watch implements grpc_health_v1.HealthServer.
func (h *healthCheckImpl) Watch(req *grpc_health_v1.HealthCheckRequest, srv grpc_health_v1.Health_WatchServer) error {
	return srv.Send(&grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	})
}

var _ grpc_health_v1.HealthServer = &healthCheckImpl{}
