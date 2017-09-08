package rpc

import (
	node_store "bitbucket.org/instinctools/gluten/master/backend/clustering"
	"bitbucket.org/instinctools/gluten/master/backend/service"
	"bitbucket.org/instinctools/gluten/shared/logging"
	pb_cli "bitbucket.org/instinctools/gluten/shared/rpc/cli"
	pb_slave "bitbucket.org/instinctools/gluten/shared/rpc/slave"
	"bitbucket.org/instinctools/gluten/shared/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
)

type RpcServer struct {
	execService *service.ExecutionService
	nodeStore   *node_store.NodeStore
	port        int
}

func NewRpcServer(execService *service.ExecutionService, nodeStore *node_store.NodeStore, port int) *RpcServer {
	var server *RpcServer
	defer server.startListening()
	server = &RpcServer{
		execService,
		nodeStore,
		port,
	}
	return server
}

func (rpc *RpcServer) startListening() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(rpc.port))
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Error during starting rpc server")
	}
	s := grpc.NewServer()
	pb_cli.RegisterProtoServiceServer(s, rpc)
	pb_slave.RegisterProtoServiceServer(s, rpc)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Error during serving")
	}
}

func (s *RpcServer) SayHello(ctx context.Context, in *pb_slave.Request) (*pb_slave.Response, error) {
	s.nodeStore.RegisterNode(in.RemoteAddress)
	return &pb_slave.Response{Message: "OK"}, nil
}

func (s *RpcServer) SendConfig(ctx context.Context, in *pb_cli.Project) (*pb_cli.ResponseMessage, error) {
	service.AddProject(utils.ParseProto2Project(in))
	s.execService.ExecuteProject(service.GetByName(in.Name))
	return &pb_cli.ResponseMessage{Message: "Project accepted & launched"}, nil
}
