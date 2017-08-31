package client

import (
	"bitbucket.org/instinctools/gluten/shared/logging"
	pb "bitbucket.org/instinctools/gluten/shared/rpc/slave"
	conf "bitbucket.org/instinctools/gluten/slave/config"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

var config *conf.Config

func init() {
	config = conf.GetConfig()
}

func StartHelloJob(address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Error while trying to connect")
	}
	defer conn.Close()
	c := pb.NewProtoServiceClient(conn)

	// Every response for server
	address = getAddress()
	for {
		_, err = c.SayHello(context.Background(), &pb.Request{RemoteAddress: address})
		if err != nil {
			logging.WithFields(logging.Fields{
				"error": err,
			}).Error("Error during sending hello message")
		}
		time.Sleep(time.Second * 5)
	}

}

func getAddress() string {
	return "192.168.0.1"
}
