package main

import (
	"context"
	"log"

	pb "github.com/altamisatmaja/commons/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server) {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Printf("Pesanan baru diterima! Order %v", p)
	o := &pb.Order{
		ID: "42",
	}

	return o, nil
}
