package LengthOfLongestSubstring

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		// Test case 1
		{
			input:    "abcabcbb",
			expected: 3,
		},
		// Test case 2
		{
			input:    "bbbbb",
			expected: 1,
		},
		// Test case 3
		{
			input:    "pwwkew",
			expected: 3,
		},
		// Additional test cases can be added here
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := lengthOfLongestSubstring(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
