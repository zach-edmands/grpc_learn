package main

import (
	"context"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc_learn/server/ecommerce"
	"log"
	"net"
)

const port = ":50051"

func main() {
	// create server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	// register the server with gRPC
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})
	log.Printf("starting gRPC listener on port %s", port)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}

// server implements the ProductInfoServer interface generated from
// the protobuffer definition by protoc.
type server struct {
	productMap map[string]*pb.Product

	// see https://github.com/grpc/grpc-go/issues/3794
	pb.UnimplementedProductInfoServer
}

func (s *server) AddProduct(ctx context.Context, product *pb.Product) (*pb.ProductID, error) {
	log.Printf("adding product: %v", product)
	product.Id = uuid.NewV4().String()

	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}

	s.productMap[product.Id] = product
	return &pb.ProductID{Value: product.Id}, nil
}

func (s *server) GetProduct(ctx context.Context, id *pb.ProductID) (*pb.Product, error) {
	log.Printf("getting product with id: %s", id.Value)
	product, ok := s.productMap[id.Value]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "product not found", id.Value)
	}
	return product, nil
}
