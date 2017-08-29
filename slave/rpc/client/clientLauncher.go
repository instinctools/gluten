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

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while trying to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProtoServiceClient(conn)

	var r *pb.Response
	var i int

	// Every response for server
	for {
		if i > 3 {
			service.ChangeStatus()
		}
		i++
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
