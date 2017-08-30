package client

import (
	"log"

	"fmt"
	"time"

	pb "bitbucket.org/instinctools/gluten/shared/rpc/slave"
	"bitbucket.org/instinctools/gluten/master/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

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
		message := service.MESSAGE
		if !service.STATUS {
			message = "405"
		}
		r, err = c.SayHello(context.Background(), &pb.Request{Message: message})
		time.Sleep(service.RESPONSE_TIME)
	}

	if err != nil {
		log.Fatalf("Error sending: %v", err)
	}
	fmt.Printf("Hello sended: %s", r.Message)
}
