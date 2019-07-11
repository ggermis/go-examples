package main

import (
	rls "github.com/envoyproxy/go-control-plane/envoy/service/ratelimit/v2"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

// Need to have an object implementing the ratelimit interface
type server struct{}

func (s *server) ShouldRateLimit(ctx context.Context, request *rls.RateLimitRequest) (*rls.RateLimitResponse, error) {
	log.Printf("request: %v, context: %v\n", request, ctx)
	status := rls.RateLimitResponse_OK
	if time.Now().Unix()%5 == 0 {
		status = rls.RateLimitResponse_OVER_LIMIT
	}
	return &rls.RateLimitResponse{OverallCode: status}, nil
}

func main() {
	// create a TCP listener on port 50052
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listening on %s", lis.Addr())

	// create a gRPC server and register the RateLimitService server
	s := grpc.NewServer()
	rls.RegisterRateLimitServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
