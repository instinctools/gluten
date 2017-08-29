package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	pb "bitbucket.org/instinctools/gluten/shared/rpc/cli"
	"golang.org/x/net/context"
)

const separator string = ":"

type server struct{}

func (s *server) SendConfig(ctx context.Context, in *pb.Project) (*pb.ResponseMessage, error) {
	fmt.Println(in)
	log.Println("Request : ", in)

	return &pb.ResponseMessage{Message: "Done"}, nil
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
		// TODO fix ports listeners 
		rpcPort := separator + *rpcCommand
		LaunchServer(rpcPort)
		webPort := separator + *webCommand
		LaunchWebServer(webPort)
		fmt.Println("Server is started")
	}
}
