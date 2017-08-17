package client

import (
	pb "bitbucket.org/instinctools/gluten/shared/rpc/slave"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func LaunchClient(address string, message string) string {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while trying to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProtoServiceClient(conn)

	r, err := c.SayHello(context.Background(), &pb.Request{message})
	if err != nil {
		log.Fatalf("Error while sending hello: %v", err)
	}
	log.Printf("Hello sended: %s", r.Message)
	return r.Message
}
