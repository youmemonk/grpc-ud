package main

import (
	"context"
	"log"
	pb "v1/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("Greet Function was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Mayank",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %v\n", res.Result)
}
