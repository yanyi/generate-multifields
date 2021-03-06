package generate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOutput(t *testing.T) {
	// Given.
	testCases := []struct {
		name      string
		startID   int
		endID     int
		inputStr  string
		expected  string
		expectErr bool
	}{
		{
			name:    "Valid input with 3 repetitions",
			startID: 1,
			endID:   3,
			inputStr: `querySomething$id: querySomething(id:$id) {
  someField
}`,
			expected: `querySomething1: querySomething(id:1) {
  someField
}
querySomething2: querySomething(id:2) {
  someField
}
querySomething3: querySomething(id:3) {
  someField
}`,
			expectErr: false,
		},
		{
			name:    "StartID and EndID is same",
			startID: 1337,
			endID:   1337,
			inputStr: `updateAnotherThing$id: updateAnotherThing(input: {
  thingId: $id
}) {
  someField
}`,
			expected: `updateAnotherThing1337: updateAnotherThing(input: {
  thingId: 1337
}) {
  someField
}`,
			expectErr: false,
		},
		{
			name:      "EndID is less than StartID",
			startID:   10,
			endID:     9,
			inputStr:  "",
			expected:  "",
			expectErr: true,
		},
		{
			name:      "Empty input string",
			startID:   4,
			endID:     5,
			inputStr:  "",
			expected:  "",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// When.
			output, err := Output(tc.startID, tc.endID, tc.inputStr)

			// Then.
			require.Equal(t, tc.expected, output)

			if tc.expectErr {
				require.Error(t, err)
			}
		})
	}
}

func ExampleOutput() {
	inputStr := `getUser$idName: getUser(id:$id) {
  name
}`
	output, _ := Output(4, 7, inputStr)

	fmt.Println(output)
	// Output:
	// getUser4Name: getUser(id:4) {
	//   name
	// }
	// getUser5Name: getUser(id:5) {
	//   name
	// }
	// getUser6Name: getUser(id:6) {
	//   name
	// }
	// getUser7Name: getUser(id:7) {
	//   name
	// }
}

func Test_validateInputs(t *testing.T) {
	// Given.
	testCases := []struct {
		name     string
		startID  int
		endID    int
		inputStr string
	}{
		{
			name:    "End should not be less than start",
			startID: 10,
			endID:   9,
		},
		{
			name: "Input string should not be empty",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// When.
			err := validateInputs(tc.startID, tc.endID, tc.inputStr)

			// Then.
			require.Error(t, err)
		})
	}
}
