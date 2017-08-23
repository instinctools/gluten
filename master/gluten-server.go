package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	pb "bitbucket.org/instinctools/gluten/shared/rpc/cli"
	"golang.org/x/net/context"
	"bitbucket.org/instinctools/gluten/master/backend"
)

const separator string = ":"

type server struct{}

func (s *server) SendConfig(ctx context.Context, in *pb.ParamsRequest) (*pb.ReplyMessage, error) {
	log.Println("Request body: ", in.Body)
	return &pb.ReplyMessage{Message: "Good day " + "sir."}, nil
}

func RunServer() {

	webCommand := flag.String("web-port", "", "port for start web-server, [:port]")
	rpcCommand := flag.String("rpc-port", "", "port for create RPC-server, [:port].")

	if len(os.Args) < 2 {
		fmt.Println("Command is wrong. Try again")
		os.Exit(1)
	}

	flag.Parse()

	if *webCommand == "" || *rpcCommand == "" {
		println(webCommand, rpcCommand)
		os.Exit(1)
	} else {
		fmt.Println("Server is started")
		webPort := separator + *webCommand
		backend.StartWebServer(webPort)
		rpcPort := separator + *rpcCommand
		LaunchServer(rpcPort)

	}
}
