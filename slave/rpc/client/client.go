package client

import (
	"bitbucket.org/instinctools/gluten/shared/logging"
	pb "bitbucket.org/instinctools/gluten/shared/rpc/slave"
	conf "bitbucket.org/instinctools/gluten/slave/config"
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
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
	address, err = getAddress()
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Error while trying to get address")
	}
	for {
		_, err = c.SayHello(context.Background(), &pb.Request{RemoteAddress: address})
		if err != nil {
			logging.WithFields(logging.Fields{
				"error": err,
			}).Error("Error during sending hello message")
		}
		time.Sleep(config.RetrieveTimeout)
	}

}

func getAddress() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			//TODO - get port from config
			return ip.String() + ":7777", nil
		}
	}
	return "", errors.New("are you connected to the network?")
}
