package server

import (
	pb "bitbucket.org/instinctools/gluten/shared/rpc/master"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
)

type server struct{}

func (s *server) SendMessage(ctx context.Context, in *pb.RequestMessage) (*pb.ResponseMessage, error) {
	log.Printf("Request message: %v", in.Message)
	return &pb.ResponseMessage{Message: in.Message}, nil
}

func LaunchServer(port int) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
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
