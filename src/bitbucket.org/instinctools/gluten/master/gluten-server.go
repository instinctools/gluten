package main

import (
	pb "bitbucket.org/instinctools/gluten/cli/proto_service"
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"log"
	"os"
	"strings"
)

const separator string = ":"

type stringList []string

func (s *stringList) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *stringList) Set(value string) error {
	*s = strings.Split(value, ",")
	return nil
}

type server struct{}

func (s *server) SendConfig(ctx context.Context, in *pb.ParamsRequest) (*pb.ReplyMessage, error) {
	log.Println("Request body: %v", in.Body)
	return &pb.ReplyMessage{Message: "Good day " + "sir."}, nil
}

func main() {

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
		web_port := separator + *webCommand
		LaunchWebServer(web_port)
		rpc_port := separator + *rpcCommand
		LaunchServer(rpc_port)

		
	}
}
