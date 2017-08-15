package main

import (
	"bitbucket.org/instinctools/gluten/slave/cmd"
	"fmt"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Info("Error at slave startup")
		os.Exit(1)
	}
}
