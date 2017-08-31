package cmd

import (
	"bitbucket.org/instinctools/gluten/master/backend/rest"
	"bitbucket.org/instinctools/gluten/master/backend/rpc"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var rpcPort, webPort string

func init() {
	RootCmd.Flags().StringVarP(&rpcPort, "rpc-port", "r", "", "Master rpc port")
	RootCmd.Flags().StringVarP(&webPort, "web-port", "w", "", "Master web port")

}

var RootCmd = &cobra.Command{
	Use:   "run",
	Short: "Same short master description",
	Long: `Longer master description.. 
            feel free to use a few lines here.
            `,
	Run: func(cmd *cobra.Command, args []string) {
		if rpcPort == "" || webPort == "" {
			os.Exit(1)
		} else {
			iRpcPort, err := strconv.ParseInt(rpcPort, 10, 64)
			if err != nil {
				panic(err)
			}
			iWebPort, err := strconv.ParseInt(webPort, 10, 64)
			if err != nil {
				panic(err)
			}
			rpc.LaunchRpcServer(int(iRpcPort))
			rest.LaunchWebServer(int(iWebPort))
		}
	},
}
