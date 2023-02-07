package main

import (
	"context"
	"log"
	pb "v2/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Sum Function was invoked with %v\n", in)

	return &pb.Response{
		Result: in.A + in.B,
	}, nil
}

func (s *Server) Difference(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Difference Function was invoked with %v\n", in)

	return &pb.Response{
		Result: in.A - in.B,
	}, nil
}

func (s *Server) Multiply(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Multiplication function was invoked with %v\n", in)

	return &pb.Response{
		Result: in.A * in.B,
	}, nil
}
