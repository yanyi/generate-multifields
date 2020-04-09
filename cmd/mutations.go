package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var mutationsCmd = &cobra.Command{
	Use:   "mutations",
	Short: "Generate multiple fields of a given mutation format",
	PreRun: func(cmd *cobra.Command, args []string) {
		// Require flags to be set before continuing with the run.
		cobra.MarkFlagRequired(cmd.Flags(), "start")
		cobra.MarkFlagRequired(cmd.Flags(), "end")
		cobra.MarkFlagRequired(cmd.Flags(), "file-path")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mutations: not implemented")
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(mutationsCmd)
}
