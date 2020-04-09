package generate

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var twoSpacesFormat = "  %v"

// FormattedInput formats the input file in preparation for generating an
// output formatted for GraphQL queries or mutations. It will add two spaces to
// the front of every line in the given input file.
func FormattedInput(file io.Reader) (string, error) {
	var strBuilder strings.Builder

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		s = fmt.Sprintf(twoSpacesFormat, s)

		fmt.Fprintln(&strBuilder, s)
	}

	// Remove "\n" in the last line.
	formattedInput := strings.TrimRight(strBuilder.String(), "\n")

	return formattedInput, nil
}
