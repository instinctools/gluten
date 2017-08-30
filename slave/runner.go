package main

import (
	"bitbucket.org/instinctools/gluten/slave/rpc/client"
)

func main() {
	client.LaunchClient(":8080")
}
