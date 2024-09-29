package main

import (
	"context"
	"fmt"
	"log"
	pb "m/protobuf/pb/permission"
	pu "m/protobuf/pb/user"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:7008")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSamplePermissionServiceServer(grpcServer, &server1{})
	pu.RegisterSampleUserServiceServer(grpcServer, &server2{})
	grpcServer.Serve(lis)
	fmt.Println("grpc server started")
}

type server1 struct {
	pb.UnimplementedSamplePermissionServiceServer
}

func (s *server1) GetPermissionDataForClient(ctx context.Context, req *pb.PermissionDataRequest) (*pb.PermissionDataResponse, error) {
	return &pb.PermissionDataResponse{Data: "Permission data for client"}, nil
}

type server2 struct {
	pu.UnimplementedSampleUserServiceServer
}

func (s *server2) GetUserDataForClient(ctx context.Context, req *pu.UserDataRequest) (*pu.UserDataResponse, error) {
	return &pu.UserDataResponse{Data: "User data for client "}, nil
}
