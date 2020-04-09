package cmd

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/yanyi/generate-multifields-gql/internal/errwrapper"
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
		err := errors.New("mutations: not implemented")
		errwrapper.Fatal(err)
	},
}

func init() {
	rootCmd.AddCommand(mutationsCmd)
}
