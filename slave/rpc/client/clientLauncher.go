package client

import (
	"log"

	"fmt"
	"time"

	pb "bitbucket.org/instinctools/gluten/shared/rpc/slave"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	conf "bitbucket.org/instinctools/gluten/slave/config"
)

var config *conf.Config

func init() {
	config = conf.GetConfig()
}

func LaunchClient(address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while trying to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProtoServiceClient(conn)

	var r *pb.Response

	// Every response for server
	for {
		message := config.Message
		r, err = c.SayHello(context.Background(), &pb.Request{Message: message})
		time.Sleep(config.RetrieveTimeout)
	}

	if err != nil {
		log.Fatalf("Error sending: %v", err)
	}
	fmt.Printf("Hello sended: %s", r.Message)
}
