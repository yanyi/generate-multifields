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
			inputStr: `createSomething(id:$id) {
  someField
}`,
			expected: `createSomething(id:1) {
  someField
}
createSomething(id:2) {
  someField
}
createSomething(id:3) {
  someField
}`,
			expectErr: false,
		},
		{
			name:    "StartID and EndID is same",
			startID: 1337,
			endID:   1337,
			inputStr: `updateAnotherThing(input: {
  thingId: $id
}) {
  someField
}`,
			expected: `updateAnotherThing(input: {
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
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
