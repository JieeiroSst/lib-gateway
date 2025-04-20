package customer_relationship_service

import (
	"context"
	"fmt"
	"time"

	pb "github.com/JIeeiroSst/lib-gateway/customer-relationship-service/gateway/customer-relationship-service"
	"github.com/JIeeiroSst/lib-gateway/pkg"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/mercari/go-circuitbreaker"
	"google.golang.org/grpc"
)

func GetClientGrpc(lnAddress string) pb.CRMServiceClient {
	time.Sleep(2 * time.Second)
	ctx := context.Background()

	cb := circuitbreaker.New(
		circuitbreaker.WithCounterResetInterval(time.Minute),
		circuitbreaker.WithTripFunc(circuitbreaker.NewTripFuncThreshold(3)),
		circuitbreaker.WithOpenTimeout(2500*time.Millisecond),
		circuitbreaker.WithHalfOpenMaxSuccesses(3),
	)

	client, err := grpc.DialContext(
		ctx,
		lnAddress,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				pkg.UnaryClientInterceptor(
					cb,
					func(ctx context.Context, method string, req interface{}) {
						fmt.Printf("[Client] Circuit breaker is open.\n")
					},
				),
			),
		),
	)
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return nil
	}
	defer client.Close()
	return pb.NewCRMServiceClient(client)
}
