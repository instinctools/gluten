package clustering

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/shared/logging"
	service "bitbucket.org/instinctools/gluten/shared/rpc/master"
	"bitbucket.org/instinctools/gluten/shared/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func SubmitOverRPC(node_address string, step core.TestStep) {
	conn, err := grpc.Dial(node_address, grpc.WithInsecure())
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Error while trying to connect")
	}
	defer conn.Close()
	c := service.NewProtoServiceClient(conn)

	_, err = c.SendMessage(context.Background(), utils.ParseStep2Proto(step))
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Error during sending hello message")
	}

}
