package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "web/protobuf_grpc/grpc/datafiles"
)

const (
	address = "localhost:50051"
)

// Client makes a single request to the GRPC server and pass details. The server picks up the details, processes them and
// sends a response back to the client.
func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewMoneyTransactionClient(conn)

	// Prepare data. Get this from clients like Frotened or App
	from := "1234"
	to := "5678"
	amount := float32(1250.75)

	// Contact the server and print out its response
	r, err := c.MakeTransaction(context.Background(), &pb.TransactionRequest{From: from, To: to, Amount: amount})
	if err != nil {
		log.Fatalf("Count not transact: %v", err)
	}
	log.Printf("Transaction confirmed: %t", r.Confirmation)
}
