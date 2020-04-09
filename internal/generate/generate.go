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

	if endID < startID {
		return "", errors.New("value of --end should not be lesser than --start")
	}

	if startID == endID {
		s := strings.Replace(input, strReplacement, strconv.Itoa(startID), -1)

		return s, nil
	}

	for id := startID; id <= endID; id++ {
		s := strings.Replace(input, strReplacement, strconv.Itoa(id), -1)

		if id == endID {
			fmt.Fprint(&outStrBuilder, s)
			break
		}
		fmt.Fprintln(&outStrBuilder, s)
	}

	return outStrBuilder.String(), nil
}
