package backend

import (
	"flag"
<<<<<<< HEAD:master/gluten-server.go
	"fmt"
	"os"

	"bitbucket.org/instinctools/gluten/master/backend"
=======
	"os"

	pb "bitbucket.org/instinctools/gluten/shared/rpc/cli"
	pu "bitbucket.org/instinctools/gluten/shared/utils"
	"golang.org/x/net/context"
>>>>>>> dev:master/backend/gluten_server.go
)

const separator string = ":"

<<<<<<< HEAD:master/gluten-server.go
=======
type server struct{}

func (s *server) SendConfig(ctx context.Context, in *pb.Project) (*pb.ResponseMessage, error) {
	log.Println("Request : ", in)
	project := pu.ParseProto2Project(in)
	log.Println("Model : ", project.GetAllSteps()[1].GetParams())

	return &pb.ResponseMessage{Message: "Done"}, nil
}

>>>>>>> dev:master/backend/gluten_server.go
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
