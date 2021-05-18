package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc_learn/client/ecommerce"
	"log"
	"time"
)

const address = "localhost:50051"

func main() {
	// connect to the gRPC server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// instantiate a client
	c := pb.NewProductInfoClient(conn)

	// add a product by calling the gRPC server
	name := "zachPhone 13"
	desc := "the best phone"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.AddProduct(
		ctx,
		&pb.Product{Name: name, Description: desc},
	)
	if err != nil {
		log.Fatalf("could not add product: %v", err)
	}
	log.Printf("product id %s added successfully", res.Value)

	// get the product back
	product, err := c.GetProduct(ctx, &pb.ProductID{Value: res.Value})
	if err != nil {
		log.Fatalf("could not get product %v", err)
	}
	log.Printf("got product: %s", product)
}
