package main

import (
	"bitbucket.org/instinctools/gluten/master/backend/cmd"
	"bitbucket.org/instinctools/gluten/shared/logging"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Info("Error at master startup")
		os.Exit(1)
	}
}
