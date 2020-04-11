// Package errwrapper provides wrappers around errors in order to not repeat
// code.
package errwrapper

import (
	"fmt"
	"os"
)

// Fatal prints the error and then exits from the application.
func Fatal(err error) {
	fmt.Println(err)
	os.Exit(1)
}
