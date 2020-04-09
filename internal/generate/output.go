// Package generate contains the logic of generating an output from a given
// input.
package generate

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	strReplacement = "$id"
)

// Output generates a given input for repeated times.
func Output(startID, endID int, input string) (string, error) {
	if err := validateInputs(startID, endID, input); err != nil {
		return "", errors.Wrapf(err, "unable to generate output")
	}

	// Handle output for only one repetition.
	if startID == endID {
		s := strings.Replace(input, strReplacement, strconv.Itoa(startID), -1)
		return s, nil
	}

	// Build the output string.
	var strBuilder strings.Builder
	for id := startID; id <= endID; id++ {
		s := strings.Replace(input, strReplacement, strconv.Itoa(id), -1)

		if id == endID {
			fmt.Fprint(&strBuilder, s)
			break
		}
		fmt.Fprintln(&strBuilder, s)
	}

	return strBuilder.String(), nil
}

func validateInputs(startID, endID int, input string) error {
	if endID < startID {
		return errors.New("value of --end should not be lesser than --start")
	}

	if len(input) == 0 || input == "" {
		return errors.New("input should not be empty")
	}

	return nil
}
