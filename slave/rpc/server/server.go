package server

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/shared/logging"
	pb "bitbucket.org/instinctools/gluten/shared/rpc/master"
	"bitbucket.org/instinctools/gluten/shared/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
)

type RpcServer struct {
	runner core.TestRunner
}

func (s *RpcServer) SendMessage(ctx context.Context, in *pb.Step) (*pb.ResponseMessage, error) {
	s.runner.Run(utils.ParseMasterProto2Step(in))
	return &pb.ResponseMessage{Message: "OK"}, nil
}

func LaunchServer(runner core.TestRunner, port int) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Info("Error during establishing connection")
	}
	s := grpc.NewServer()
	pb.RegisterProtoServiceServer(s, &RpcServer{runner})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Info("Error during serving incomming request")
	}
}
