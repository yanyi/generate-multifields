package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/yanyi/generate-multifields/internal/errwrapper"
	"github.com/yanyi/generate-multifields/internal/generate"
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
	Run: runFunc,
}

func init() {
	rootCmd.AddCommand(mutationsCmd)
}

// runFunc is used by `cobra.Command.Run`. It does the bulk of the job.
func runFunc(cmd *cobra.Command, args []string) {
	fileStr, err := readFile(FormatFilePath)
	if err != nil {
		errwrapper.Fatal(err)
	}

	output, err := generate.Output(StartID, EndID, fileStr)
	if err != nil {
		err := errors.Wrapf(err, "unable to generate output")
		errwrapper.Fatal(err)
	}

	// Append 'mutation { }' to wrap around given output.
	var result bytes.Buffer
	fmt.Fprintln(&result, "mutation {")
	fmt.Fprintln(&result, output)
	fmt.Fprint(&result, "}")

	// Print the result to os.Stdout.
	fmt.Print(result.String())
}

// readFile opens the given file at the file path and formats the given input
// before returning.
func readFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", errors.Wrapf(err, "unable to open file")
	}

	formattedInput, err := generate.FormattedInput(file)
	if err != nil {
		return "", errors.Wrapf(err, "unable to format file")
	}
	file.Close()

	return formattedInput, nil
}
