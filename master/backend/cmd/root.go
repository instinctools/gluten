package cmd

import (
	"bitbucket.org/instinctools/gluten/master/backend/rest"
	"bitbucket.org/instinctools/gluten/master/backend/rpc"
	"github.com/spf13/cobra"
	"os"
)

var rpcPort, webPort int

func init() {
	RootCmd.Flags().IntVarP(&rpcPort, "rpc-port", "r", 0, "Master rpc port")
	RootCmd.Flags().IntVarP(&webPort, "web-port", "w", 0, "Master web port")
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
			go rpc.LaunchRpcServer(rpcPort)
			rest.LaunchWebServer(webPort)
		}
	},
}
