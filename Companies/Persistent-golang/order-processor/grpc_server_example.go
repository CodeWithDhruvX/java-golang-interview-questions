package main

// NOTE: This file is for demonstration purposes.
// It requires code generation via protoc to correctly compile if uncommented.
// To generate:
// protoc --go_out=. --go_grpc_out=. proto/order.proto

/*
import (
	"context"
	"log"

	pb "order-processor/proto" // API assumed to be generated here
)

// GRPCServer implements the generated OrderServiceServer interface
type GRPCServer struct {
	pb.UnimplementedOrderServiceServer
	JobQueue chan Job
}

func (s *GRPCServer) SubmitOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	log.Printf("Received gRPC Order: %s", req.Id)

	// Reuse the same logic as HTTP handler
	select {
	case s.JobQueue <- Job{Order: Order{ID: req.Id, Value: req.Value}}:
		return &pb.OrderResponse{
			Id:      req.Id,
			Status:  "Queued",
			Success: true,
		}, nil
	default:
		return &pb.OrderResponse{
			Id:      req.Id,
			Status:  "Queue Full",
			Success: false,
		}, nil
	}
}
*/
