package main

import (
	rpc "bitbucket.org/instinctools/gluten/master/backend/rpc"
)

const separator string = ":"

func main() {
	rpc.LaunchServer(8888)
}
