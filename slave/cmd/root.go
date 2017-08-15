package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"bitbucket.org/instinctools/gluten/slave/rpc/client"
)

var masterIp, slavePort string

var RootCmd = &cobra.Command{
	Use:   "gluten-slave",
	Short: "Short description",
	Long: `Longer description.. 
            feel free to use a few lines here.
            `,
	Run: func(cmd *cobra.Command, args []string) {
		response := client.LaunchClient(masterIp, slavePort)
		fmt.Println(response)
	},
}

func init() {
    RootCmd.PersistentFlags().StringVar(&masterIp, "master", "", "Master IP")
    RootCmd.PersistentFlags().StringVar(&slavePort, "rpc-port", "", "Slave port")
}