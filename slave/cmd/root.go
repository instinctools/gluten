package cmd

import (
	//	"bitbucket.org/instinctools/gluten/slave/rpc/client"
	"github.com/spf13/cobra"
)

var masterIp, slavePort string

func init() {
	RootCmd.Flags().StringVarP(&masterIp, "master", "m", "", "Master IP")
	RootCmd.Flags().StringVarP(&slavePort, "rpc-port", "rp", "", "Slave port")
}

var RootCmd = &cobra.Command{
	Use:   "gluten-slave",
	Short: "Same short gluten description",
	Long: `Longer gluten description.. 
            feel free to use a few lines here.
            `,
	Run: func(cmd *cobra.Command, args []string) {

		//TODO - what is it?
		//		client.LaunchClient(masterIp, slavePort)
	},
}
