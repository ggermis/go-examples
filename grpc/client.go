package main

import (
	rls "github.com/envoyproxy/go-control-plane/envoy/service/ratelimit/v2"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func doRequest(c rls.RateLimitServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ShouldRateLimit(ctx, &rls.RateLimitRequest{Domain: "envoy"})
	if err != nil {
		log.Fatalf("could not call service: %v", err)
	}

	log.Printf("response: %v", r)
}

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	c := rls.NewRateLimitServiceClient(conn)

	// Send requests to the server
	for i := 1; i <= 13; i++ {
		doRequest(c)
		time.Sleep(1 * time.Second)
	}
}
