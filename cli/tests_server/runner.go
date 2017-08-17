package main

import (
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "bitbucket.org/instinctools/gluten/shared/rpc/cli"
	"google.golang.org/grpc/reflection"
)

var address string = ":3000"

type server struct{}

func (s *server) SendConfig(ctx context.Context, in *pb.ParamsRequest) (*pb.ReplyMessage, error) {
	log.Println("Request body: %v", in.Body)
	return &pb.ReplyMessage{Message: "Good day " + "sir."}, nil
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProtoServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
