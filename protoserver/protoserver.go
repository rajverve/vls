package protoserver

import (
	pb "github.com/rajverve/protobuf"
	"fmt"
	"golang.org/x/net/context"
)
type ProtoServer struct {

}

func (p ProtoServer) GetSegmentInfo (c context.Context, req *pb.AdRequest) (*pb.SupplySegment, error) {
	fmt.Printf("Received an AdRequest %v\n", req)
	fmt.Printf("Context %v\n", c)

	return &pb.SupplySegment{
		Audience: "Soccer Mom",
		PlaceName: "Target Encinitas",
		PlaceType: "Shopping Center",
	}, nil
}