package cmd

import (
	"os"

	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/core/result_handlers"
	"bitbucket.org/instinctools/gluten/master/backend/config"
	"bitbucket.org/instinctools/gluten/master/backend/migration"
	"bitbucket.org/instinctools/gluten/master/backend/rest"
	"bitbucket.org/instinctools/gluten/master/backend/rpc"
	"bitbucket.org/instinctools/gluten/master/backend/service"
	"bitbucket.org/instinctools/gluten/shared/logging"
	"bitbucket.org/instinctools/gluten/shared/persistence/gorm"
	"github.com/spf13/cobra"
)

var rpcPort, webPort int

func init() {
	RootCmd.Flags().IntVarP(&rpcPort, "rpc-port", "r", 0, "Master rpc port")
	RootCmd.Flags().IntVarP(&webPort, "web-port", "w", 0, "Master web port")
	migrationsConfig := config.GlobalConfig.DB.Migrations
	if migrationsConfig.Enable {
		migration.ApplyMigrations(migrationsConfig.Folder, migrationsConfig.ConnectionString)
	}
}

var RootCmd = &cobra.Command{
	Use:   "run",
	Short: "Same short master description",
	Long: `Longer master description.. 
            feel free to use a few lines here.
            `,
	Run: func(cmd *cobra.Command, args []string) {
		if rpcPort == 0 || webPort == 0 {
			os.Exit(1)
		} else {
			go rpc.LaunchRpcServer(service.ExecutionServiceInstance, rpcPort)
			logging.WithFields(logging.Fields{"port": rpcPort}).Error("Rpc server has been successfully started on port")
			rest.LaunchWebServer(webPort)
		}
	},
}
