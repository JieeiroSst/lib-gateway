package shipping_service

import (
	"context"
	"fmt"
	"time"

	"github.com/JIeeiroSst/lib-gateway/pkg"
	pb "github.com/JIeeiroSst/lib-gateway/shipping-service/gateway/shipping-service"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/mercari/go-circuitbreaker"
	"google.golang.org/grpc"
)

func GetClientGrpc(lnAddress string) pb.ShippingServiceClient {
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
	return pb.NewShippingServiceClient(client)
}
