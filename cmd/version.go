package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "v0.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of generate-multifields",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate-multifields", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
