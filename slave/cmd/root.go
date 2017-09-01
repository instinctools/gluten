package cmd

import (
	"bitbucket.org/instinctools/gluten/slave/rpc/client"
	"bitbucket.org/instinctools/gluten/slave/rpc/server"
	"github.com/spf13/cobra"
)

var (
	masterIp  string
	slavePort int
)

func init() {
	RootCmd.Flags().StringVarP(&masterIp, "master", "m", "", "Master IP")
	RootCmd.Flags().IntVarP(&slavePort, "rpc-port", "r", 0, "Slave port")
}

var RootCmd = &cobra.Command{
	Use:   "gluten-slave",
	Short: "Same short gluten description",
	Long: `Longer gluten description.. 
            feel free to use a few lines here.
            `,
	Run: func(cmd *cobra.Command, args []string) {
		go client.StartHelloJob(masterIp)
		server.LaunchServer(slavePort)
	},
}
