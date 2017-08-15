package client

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "bitbucket.org/instinctools/gluten/cli/proto_service"
)

func LaunchClient(address string, json string)  {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProtoServiceClient(conn)


	r, err := c.SendConfig(context.Background(), &pb.ParamsRequest{Url: address, Body: json})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Configurate: %s", r.Message)
}
