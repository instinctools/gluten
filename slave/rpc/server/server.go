package main

import (
	pb "bitbucket.org/instinctools/gluten/shared/rpc/master"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var address string = ":3000"

type server struct{}

func (s *server) SendMessage(ctx context.Context, in *pb.RequestMessage) (*pb.ResponseMessage, error) {
	log.Printf("Request message: %v", in.Message)
	return &pb.ResponseMessage{Message: in.Message}, nil
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Listen error: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProtoServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Serve error: %v", err)
	}
}
