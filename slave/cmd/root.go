package cmd

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/core/result_handlers"
	"bitbucket.org/instinctools/gluten/shared/logging"
	"bitbucket.org/instinctools/gluten/shared/persistence/gorm"
	"bitbucket.org/instinctools/gluten/slave/rpc/client"
	"bitbucket.org/instinctools/gluten/slave/rpc/server"
	"github.com/spf13/cobra"
)

var (
	masterIp string
	rpcPort  int
)

func init() {
	RootCmd.Flags().StringVarP(&masterIp, "master", "m", "", "Master IP")
	RootCmd.Flags().IntVarP(&rpcPort, "rpc-port", "r", 0, "Rpc port")
}

var RootCmd = &cobra.Command{
	Use:   "gluten-slave",
	Short: "Same short gluten description",
	Long: `Longer gluten description.. 
            feel free to use a few lines here.
            `,
	Run: func(cmd *cobra.Command, args []string) {
		runner := core.NewDefaultRunner(result_handlers.NewPersistentResultHandler(gorm.NewGormResultsRepo()))
		logging.WithFields(logging.Fields{"port": rpcPort}).Error("Rpc server has been successfully started on port")
		go server.LaunchServer(runner, rpcPort)
		client.StartHelloJob(masterIp)

	},
}
