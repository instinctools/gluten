package cmd

import (
	"bitbucket.org/instinctools/gluten/cli/client"
	"github.com/spf13/cobra"
	"os"
)

var pathToGenFile string

func init() {
	// TODO: add default path 
	generateCmd.Flags().StringVarP(&pathToGenFile, "path-to-file", "p", "", "Path to file 4 gen json")
	RootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate json to file",
	Long:  `Long description`,
	Run: func(cmd *cobra.Command, args []string) {
		if pathToGenFile == "" {
			os.Exit(1)
		} else {
			client.AutoGenerateConfig(pathToGenFile)
			os.Exit(1)
		}
	},
}
