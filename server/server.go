package main

import (
	pb "48HW/proto"
	"context"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStringServiceServer
}

func (s *server) ToUpperCase(ctx context.Context, in *pb.StringRequest) (*pb.StringReply, error) {
	return &pb.StringReply{Result: strings.ToUpper(in.Input)}, nil
}

func (s *server) ToLowerCase(ctx context.Context, in *pb.StringRequest) (*pb.StringReply, error) {
	return &pb.StringReply{Result: strings.ToLower(in.Input)}, nil
}

func (s *server) ReverseString(ctx context.Context, in *pb.StringRequest) (*pb.StringReply, error) {
	runes := []rune(in.Input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return &pb.StringReply{Result: string(runes)}, nil
}

func (s *server) GetStringLength(ctx context.Context, in *pb.StringRequest) (*pb.LengthReply, error) {
	return &pb.LengthReply{Length: int32(len(in.Input))}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStringServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
