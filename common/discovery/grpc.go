package discovery

import (
	"context"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"log"
	"math/rand"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ServiceConnection(ctx context.Context, serviceName string, serviceRegistry ServiceRegistry) (*grpc.ClientConn, error) {
	serviceAddress, err := serviceRegistry.Discover(ctx, serviceName)
	if err != nil {
		return nil, err
	}

	log.Printf("Discovered %d instances of %s", len(serviceAddress), serviceName)

	// Randomly select an instance
	return grpc.NewClient(
		serviceAddress[rand.Intn(len(serviceAddress))],
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
}
