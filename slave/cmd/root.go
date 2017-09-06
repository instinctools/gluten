package cmd

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/core/result_handlers"
	"bitbucket.org/instinctools/gluten/shared/logging"
	"bitbucket.org/instinctools/gluten/shared/persistence/gorm"
	"bitbucket.org/instinctools/gluten/slave/config"
	"bitbucket.org/instinctools/gluten/slave/rpc/client"
	"bitbucket.org/instinctools/gluten/slave/rpc/server"
	"github.com/spf13/cobra"
)

var (
	conf    *config.Config
	master  string
	rpcPort int
)

func init() {
	conf = config.GolbalConfig
	RootCmd.Flags().StringVarP(&master, "master", "m", "", "Master Url")
	RootCmd.Flags().IntVarP(&rpcPort, "rpc-port", "r", 0, "Rpc port")
}

func overrideConfig() {
	if master != "" {
		conf.MasterUrl = master
	}
	if rpcPort > 0 {
		conf.RpcPort = rpcPort
	}

}

var RootCmd = &cobra.Command{
	Use:   "gluten-slave",
	Short: "Same short gluten description",
	Long: `Longer gluten description.. 
            feel free to use a few lines here.
            `,
	Run: func(cmd *cobra.Command, args []string) {
		overrideConfig()
		runner := core.NewDefaultRunner(result_handlers.NewPersistentResultHandler(gorm.NewGormResultsRepo(conf.DBUrl)))
		logging.WithFields(logging.Fields{"port": conf.RpcPort}).Error("Rpc server has been successfully started on port")
		go server.LaunchServer(runner, conf.RpcPort)
		client.StartHelloJob(conf.MasterUrl, conf.RetrieveTimeout, conf.RpcPort)

	},
}
