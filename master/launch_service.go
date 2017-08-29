package main

import (
	"log"
	"net"
	"bitbucket.org/instinctools/gluten/master/service"
	pb_cli "bitbucket.org/instinctools/gluten/shared/rpc/cli"
	pb_slave "bitbucket.org/instinctools/gluten/shared/rpc/slave"

	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc/reflection"
)

var client_address string

func (s *server) SayHello(ctx context.Context, in *pb_slave.Request) (*pb_slave.Response, error) {
	CheckRequest(in.Message)
	return &pb_slave.Response{Message: in.Message}, nil
}

func LaunchServer(address string) {
	lis, err := net.Listen("tcp", address)
	DefineClientAddress(lis)
	if err != nil {
		log.Fatalf("Listen error: %v", err)
	}
	s := grpc.NewServer()
	pb_cli.RegisterProtoServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Serve error: %v", err)
	}
}

func DefineClientAddress(lis net.Listener) {
	conn, err := lis.Accept()
	if err != nil {
		log.Fatal("Connection error")
	}
	client_address = conn.RemoteAddr().String()
	conn.Close()
}

func CheckStatusSlave(in string) bool {
	answer := true
	if in != "200" {
		answer = false
	}
	return answer
}

func CheckRequest(in string) {
	status := CheckStatusSlave(in)
	if status {
		service.AddNode(client_address)
	} else {
		service.AddResponseTimeForNode(client_address)
		if service.Exist(client_address) && service.GetByKey(client_address) >= service.EXIT_TIME {
			service.RemoveNode(client_address)
		}
	}
}
