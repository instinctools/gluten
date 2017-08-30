package main

import (
	"bitbucket.org/instinctools/gluten/slave/cmd"
	"os"
	log "bitbucket.org/instinctools/gluten/shared/logging"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Info("Error at slave startup")
		os.Exit(1)
	}
}
