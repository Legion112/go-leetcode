package AddTwoNumbers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	testCases := []struct {
		name     string
		first    *ListNode
		second   *ListNode
		expected *ListNode
	}{
		{
			name:     "342+465=708",
			first:    &ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			second:   &ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			expected: &ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
		{
			name:     "0+0=0",
			first:    &ListNode{0, nil},
			second:   &ListNode{0, nil},
			expected: &ListNode{0, nil},
		},
		{
			name:     "9999999+9999=89990001",
			first:    &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}},
			second:   &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}},
			expected: &ListNode{8, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}},
		},
		{
			name:     "0+12=0",
			first:    &ListNode{2, &ListNode{1, nil}},
			second:   &ListNode{0, nil},
			expected: &ListNode{2, &ListNode{1, nil}},
		},
		{
			name:     "12+0=0",
			first:    &ListNode{0, nil},
			second:   &ListNode{2, &ListNode{1, nil}},
			expected: &ListNode{2, &ListNode{1, nil}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := addTwoNumbers(tc.first, tc.second)
			require.Equal(t, tc.expected, result, "Structs should be equal expected: %v got: %v", tc.expected, result)
		})
	}

}
