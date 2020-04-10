package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version is the CLI's version. It can be replaced with `-ldflags="-X"`.
	Version = "Development"
	// Build is the Git SHA. Also replaceable with `-ldflags="-X"`.
	Build = "4c17aadf5117487aab7bc50cbf056caf3977cc31"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of generate-multifields",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate-multifields:")
		fmt.Println(" Version:", Version)
		fmt.Println(" Build:", Build)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
