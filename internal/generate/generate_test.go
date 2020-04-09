package generate

import (
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
			inputStr: `createSomething$id: createSomething(id:$id) {
  someField
}`,
			expected: `createSomething1: createSomething(id:1) {
  someField
}
createSomething2: createSomething(id:2) {
  someField
}
createSomething3: createSomething(id:3) {
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
