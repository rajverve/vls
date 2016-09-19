package main

import (
	"net"
	"fmt"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc"
	"github.com/rajverve/vls/protoserver"
	pb "github.com/rajverve/protobuf"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 4444))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("Starting VLS server...")
	grpcServer := grpc.NewServer()
	pb.RegisterSegmentationServer(grpcServer, protoserver.ProtoServer{})
	grpcServer.Serve(lis)
}
