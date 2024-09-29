package main

import (
	"context"
	"fmt"
	"log"
	pb "m/protobuf/pb/permission"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:7008", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at localhost:50051: %v", err)
	}
	defer conn.Close()
	c := pb.NewSamplePermissionServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetPermissionDataForClient(ctx, &pb.PermissionDataRequest{})
	if err != nil {
		log.Fatalf("error calling function SayHello: %v", err)
	}
	fmt.Println(r)
}
