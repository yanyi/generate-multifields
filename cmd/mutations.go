package cmd

import (
	"errors"
	"strings"

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

		if EndID < StartID {
			err := errors.New("value of --end should not be lesser than --start")
			errwrapper.Fatal(err)
		}

		filePath := strings.TrimSpace(FormatFilePath)
		if len(filePath) == 0 || filePath == "" {
			err := errors.New("value of --file-path should not be empty")
			errwrapper.Fatal(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := errors.New("mutations: not implemented")
		errwrapper.Fatal(err)
	},
}

func init() {
	rootCmd.AddCommand(mutationsCmd)
}
