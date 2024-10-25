package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/gkwa/everymole/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of everymole",
	Long:  `All software has versions. This is everymole's`,
	Run: func(cmd *cobra.Command, args []string) {
		buildInfo := version.GetBuildInfo()
		fmt.Println(buildInfo)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
