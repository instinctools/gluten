package backend

import (
	"log"
	"net"

	pb "bitbucket.org/instinctools/gluten/shared/rpc/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func LaunchServer(address string) {
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
