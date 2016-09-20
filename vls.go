package main

import (
	"fmt"
	pb "github.com/rajverve/protobuf"
	"github.com/rajverve/vls/protoserver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 4444))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
		return
	}

	fmt.Println("Starting VLS server...")
	grpcServer := grpc.NewServer()
	pb.RegisterSegmentationServer(grpcServer, protoserver.ProtoServer{})
	grpcServer.Serve(lis)
}
