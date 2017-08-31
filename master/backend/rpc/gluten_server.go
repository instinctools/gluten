package rpc

import (
	store "bitbucket.org/instinctools/gluten/master/backend/nodestore"
	"bitbucket.org/instinctools/gluten/shared/logging"
	pb_cli "bitbucket.org/instinctools/gluten/shared/rpc/cli"
	pb_slave "bitbucket.org/instinctools/gluten/shared/rpc/slave"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
)

type server struct{}

func LaunchServer(port int) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Error during starting rpc server")
	}
	s := grpc.NewServer()
	pb_cli.RegisterProtoServiceServer(s, &server{})
	pb_slave.RegisterProtoServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Error during serving")
	}
}

func (s *server) SayHello(ctx context.Context, in *pb_slave.Request) (*pb_slave.Response, error) {
	store.RegisterNode(in.Message, ResolveSlaveAddress())
	return &pb_slave.Response{Message: in.Message}, nil
}

func (s *server) SendConfig(ctx context.Context, in *pb_cli.Project) (*pb_cli.ResponseMessage, error) {
	return &pb_cli.ResponseMessage{Message: "Good day " + "sir."}, nil
}

func ResolveSlaveAddress() string {
	/*	conn, err := lis.Accept()
		defer conn.Close()
		if err != nil {
			logging.WithFields(logging.Fields{
				"error": err,
			}).Error("Slave address can't be resolved")
		}
		return conn.RemoteAddr().String()
	*/
	return "127.0.0.1"
}
