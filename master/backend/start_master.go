package main

import (
	rpc "bitbucket.org/instinctools/gluten/master/backend/rpc"
)

func main() {
	rpc.LaunchRpcServer(8888)
}
