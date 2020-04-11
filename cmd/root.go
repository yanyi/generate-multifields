package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "generate-multifields",
		Short: "Generate multiple fields for GraphQL queries or mutations",
		Long: `generate-multifields is a CLI library to help generate multiple
fields of a given input format of your GraphQL queries or mutations, by
repeating them for a number of times.`,
	}

	// StartID is the starting ID number to start with.
	StartID int
	// EndID is the ending ID number to end with.
	EndID int
	// FormatFilePath is the file path to look for the file format in order to
	// generate repetitions and string replacement of it.
	FormatFilePath string
)

func init() {
	// Flags.
	pf := rootCmd.PersistentFlags()
	pf.IntVarP(&StartID, "start", "s", 0, "id number to start with (required)")
	pf.IntVarP(&EndID, "end", "e", 1, "id number to start with (required)")
	pf.StringVarP(&FormatFilePath, "file-path", "f", "", "path to the file format for generating repetitions (required)")
}

// Execute is root command's execution.
func Execute() error {
	return rootCmd.Execute()
}
