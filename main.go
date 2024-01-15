package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "golang-gRPC-example/proto"
	// "golang-gRPC-example/services"

	"google.golang.org/grpc"
)

type productServiceServer struct {
	pb.UnimplementedProductServiceServer
}

var products = map[int32]*pb.ProductResponse{
	1: {Id: 1, Name: "Product 1", Price: 100},
	2: {Id: 2, Name: "Product 2", Price: 150},
}

func (s *productServiceServer) GetProductByID(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	product, ok := products[req.Id]
	if !ok {
		return nil, fmt.Errorf("Product not found")
	}
	return product, nil
}

func main() {
	lis, err := net.Listen("tcp", "[::1]:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	service := &productServiceServer{}

	pb.RegisterProductServiceServer(grpcServer, service)
	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalf("Error strating server: %v", err)
	}
}
