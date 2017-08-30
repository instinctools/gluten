package client

import (
	"bitbucket.org/instinctools/gluten/shared/logging"
	pb "bitbucket.org/instinctools/gluten/shared/rpc/cli"
	pu "bitbucket.org/instinctools/gluten/shared/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func SendJsonToMaster(address string, json string) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Can't connect to master")
	}
	defer conn.Close()
	c := pb.NewProtoServiceClient(conn)
	message := pu.DeserializeJsonToProto(json)
	r, err := c.SendConfig(context.Background(), message)
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Master throws error")
	}
	logging.WithFields(logging.Fields{
		"message": r.Message,
	}).Info("Master accept configuration")
}
