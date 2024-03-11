package operator_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func TestHealth(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:5055", grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	defer conn.Close()

	client := grpc_health_v1.NewHealthClient(conn)
	response, err := client.Check(context.TODO(), &grpc_health_v1.HealthCheckRequest{})
	require.NoError(t, err)
	require.Equal(t, grpc_health_v1.HealthCheckResponse_SERVING, response.Status)

}
