package main

import (
	"bitbucket.org/instinctools/gluten/cli/cmd"
	"bitbucket.org/instinctools/gluten/shared/logging"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Can't execute RootCmd")
		os.Exit(1)
	}
}
