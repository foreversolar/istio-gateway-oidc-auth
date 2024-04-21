package main

import (
	"context"
	"log"
	"time"

	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func _main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := auth.NewAuthorizationClient(conn)

	testCases := []struct {
		description  string
		token        string
		expectedCode int32
	}{
		{"Valid token", "Bearer foo", 0}, // rpc.OK
		{"Invalid token", "Be", 16},      // rpc.PERMISSION_DENIED
		{"Missing token", "", 16},        // rpc.UNAUTHENTICATED
	}

	for index, tc := range testCases {
		req := &auth.CheckRequest{
			Attributes: &auth.AttributeContext{
				Request: &auth.AttributeContext_Request{
					Http: &auth.AttributeContext_HttpRequest{
						Headers: map[string]string{"authorization": tc.token},
					},
				},
			},
		}

		if index == 2 {
			req = &auth.CheckRequest{
				Attributes: &auth.AttributeContext{
					Request: &auth.AttributeContext_Request{
						Http: &auth.AttributeContext_HttpRequest{
							Headers: map[string]string{},
						},
					},
				},
			}
		}

		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		resp, err := c.Check(ctx, req)
		if err != nil {
			log.Fatalf(err.Error())
		}
		if resp.Status.Code != tc.expectedCode {
			log.Fatalf("%s: Expected status code %v, got %v", tc.description, tc.expectedCode, resp.Status.Code)
		}
	}
	log.Println("All tests passed.")
}
