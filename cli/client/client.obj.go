package client

import (
	"log"
	pb "bitbucket.org/instinctools/gluten/shared/rpc/cli"
	des "bitbucket.org/instinctools/gluten/shared/deserializers"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func LaunchClient(address string, json string) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProtoServiceClient(conn)
	message := des.DeserializeJsonToProto(json)
	r, err := c.SendConfig(context.Background(), message)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Configurate: %s", r.Message)
}
