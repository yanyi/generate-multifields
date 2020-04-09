// Package generate contains the logic of generating an output from a given
// input.
package generate

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	strReplacement = "$id"
)

// Output generates a given input for repeated times.
func Output(startID, endID int, input string) (string, error) {
	var outStrBuilder strings.Builder

	if startID == endID {
		return "", errors.New("not implemented")
	}

	for id := startID; id <= endID; id++ {
		// Take input and replace strReplacement with id.
		s := strings.Replace(input, strReplacement, strconv.Itoa(id), -1)

		if id == endID {
			fmt.Fprint(&outStrBuilder, s)
			break
		}
		fmt.Fprintln(&outStrBuilder, s)
	}

	return outStrBuilder.String(), nil
}
