package ValidateParantheses

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		input  string
		output bool
	}{
		{
			input:  "()",
			output: true,
		},
		{
			input:  "()[]{}",
			output: true,
		},
		{
			input:  "(]",
			output: false,
		},
		{
			"[", false,
		},
	}

	for _, test := range testCases {
		t.Run(test.input, func(t *testing.T) {
			require.Equal(t, test.output, isValid(test.input))
		})
	}
}
