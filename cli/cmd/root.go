package cmd

import (
	"github.com/spf13/cobra"
	"bitbucket.org/instinctools/gluten/cli/client"
	"os"
)

var masterPort, pathToJsonFile string

func init() {
    RootCmd.Flags().StringVarP(&masterPort, "master-port", "m", "", "Master port")
    RootCmd.Flags().StringVarP(&pathToJsonFile, "path-to-file", "p", "", "Path to json file")
    
}

var RootCmd = &cobra.Command{
	Use:   "run",
	Short: "Same short cli description",
	Long: `Longer cli description.. 
            feel free to use a few lines here.
            `,
	Run: func(cmd *cobra.Command, args []string) {
		if masterPort == "" || pathToJsonFile == "" {
			os.Exit(1)
		} else {
			json := client.ReadJSONFile(pathToJsonFile)
			client.SendJsonToMaster(masterPort, json)
			os.Exit(1)
		}
	},
}
