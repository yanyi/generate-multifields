package cmd

import (
	"fmt"
	"io"
	"runtime"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var (
	// Version is the CLI's version. It can be replaced with `-ldflags="-X"`.
	Version = "Development"
	// Build is the Git SHA. Also replaceable with `-ldflags="-X"`.
	Build = "4c17aadf5117487aab7bc50cbf056caf3977cc31"

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of generate-multifields",
		Run: func(cmd *cobra.Command, args []string) {
			prettyPrint(cmd.OutOrStdout())
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

func prettyPrint(out io.Writer) {
	fmt.Println("generate-multifields:")

	w := tabwriter.NewWriter(out, 10, 1, 1, ' ', 0)
	defer w.Flush()

	fmt.Fprintln(w, " Version:\t", Version)
	fmt.Fprintln(w, " Build:\t", Build)
	fmt.Fprintln(w, " OS/Arch:\t", fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))
}
