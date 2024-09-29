package main

import (
	"context"
	"fmt"
	"log"
	pu "m/protobuf/pb/user"
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
	c := pu.NewSampleUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetUserDataForClient(ctx, &pu.UserDataRequest{})
	if err != nil {
		log.Fatalf("error calling function SayHello: %v", err)
	}
	fmt.Println(r)

}
