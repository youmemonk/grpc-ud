package main

import (
	"context"
	"log"
	pb "v2/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient, a uint32, b uint32) {
	log.Printf("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.Request{
		A: a,
		B: b,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Sum: %v\n", res.Result)
}

func doDifference(c pb.CalculatorServiceClient, a uint32, b uint32) {
	log.Printf("doDifference was invoked")
	res, err := c.Difference(context.Background(), &pb.Request{
		A: a,
		B: b,
	})

	if err != nil {
		log.Fatalf("Could not subt: %v\n", err)
	}

	log.Printf("Difference: %v\n", res.Result)
}

func doMultiply(c pb.CalculatorServiceClient, a uint32, b uint32) {
	log.Printf("doMultiply was invoked")
	res, err := c.Multiply(context.Background(), &pb.Request{
		A: a,
		B: b,
	})

	if err != nil {
		log.Fatalf("Could not Multiply: %v\n", err)
	}

	log.Printf("Multiplication: %v\n", res.Result)
}
