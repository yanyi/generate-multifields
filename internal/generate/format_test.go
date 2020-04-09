package generate

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormattedInput(t *testing.T) {
	// Given.
	fileContent := `querySomething$id: something(id: $id) {
  someField
}`
	expected := `  querySomething$id: something(id: $id) {
    someField
  }`

	// When.
	formattedInput, _ := FormattedInput(strings.NewReader(fileContent))

	// Then.
	require.Equal(t, expected, formattedInput)
}

func ExampleFormattedInput() {
	fileContent := `queryContainer$id: container(id: $id) {
  name
}`

	formattedInput, _ := FormattedInput(strings.NewReader(fileContent))
	fmt.Println(formattedInput)
	// Output:
	//   queryContainer$id: container(id: $id) {
	//     name
	//   }
}
