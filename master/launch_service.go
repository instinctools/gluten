package main

import (
	"log"
	"net"

	"bitbucket.org/instinctools/gluten/master/service"
	pb_cli "bitbucket.org/instinctools/gluten/shared/rpc/cli"
	pb_master "bitbucket.org/instinctools/gluten/shared/rpc/master"
	pb_slave "bitbucket.org/instinctools/gluten/shared/rpc/slave"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"time"
)

var client_address string

type server struct{}

func (s *server) SendConfig(ctx context.Context, in *pb_cli.ParamsRequest) (*pb_cli.ReplyMessage, error) {
	log.Println("Request body: ", in.Body)
	return &pb_cli.ReplyMessage{Message: "Good day " + "sir."}, nil
}

func (s *server) SayHello(ctx context.Context, in *pb_slave.Request) (*pb_slave.Response, error) {
	CheckRequest(in.Message)
	return &pb_slave.Response{Message: in.Message}, nil
}

func LaunchServer(address string) {
	lis, err := net.Listen("tcp", address)
	DefineSlaveAddress(lis)
	if err != nil {
		log.Fatalf("Listen error: %v", err)
	}
	s := grpc.NewServer()
	pb_cli.RegisterProtoServiceServer(s, &server{})
	pb_slave.RegisterProtoServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Serve error: %v", err)
	}
}

func LaunchClient(address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb_master.NewProtoServiceClient(conn)
	r, err := c.SendMessage(context.Background(), &pb_master.RequestMessage{Message: "no implementation"})
	if err != nil {
		log.Fatalf("not response: %v", err)
	}
	log.Printf("Testing message: %s", r.Message)
}

func DefineSlaveAddress(lis net.Listener) {
	conn, err := lis.Accept()
	if err != nil {
		log.Fatal("Connection error")
	}
	client_address = conn.RemoteAddr().String()
	conn.Close()
}

func CheckSlaveStatus(in string) bool {
	answer := true
	if in != service.MESSAGE {
		answer = false
	}
	return answer
}

func CheckRequest(in string) {
	status := CheckSlaveStatus(in)
	if status {
		service.AddNode(client_address)
	} else {
		service.AddResponseTimeForNode(client_address)
		if service.Exist(client_address) && service.GetByKey(client_address) >= service.EXIT_TIME {
			service.RemoveNode(client_address)
		}
	}
}

func main() {
	nanos := time.Now().UnixNano()
	ti := nanos - time.Hour.Nanoseconds()
	println(time.Unix(0, nanos).String())
	print(time.Unix(0, ti).String())
}
