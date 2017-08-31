package main

import (
	//	"bitbucket.org/instinctools/gluten/shared/logging"
	//	"bitbucket.org/instinctools/gluten/slave/cmd"
	"bitbucket.org/instinctools/gluten/slave/rpc/client"
	//	"os"
)

func main() {
	client.StartHelloJob("localhost:8888")
	/*	if err := cmd.RootCmd.Execute(); err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Info("Error at slave startup")
		os.Exit(1)
	}*/
}
