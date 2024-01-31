package AddTwoNumbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	testCases := []struct {
		first    *ListNode
		second   *ListNode
		expected *ListNode
	}{
		{
			first:    &ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			second:   &ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			expected: &ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
		{
			first:    &ListNode{0, nil},
			second:   &ListNode{0, nil},
			expected: &ListNode{0, nil},
		},
		{
			first:    &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}},
			second:   &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}},
			expected: &ListNode{8, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}},
		},
	}

	for index, tc := range testCases {
		result := addTwoNumbers(tc.first, tc.second)
		assert.Equal(t, tc.expected, result, "Index: %d Structs should be equal expected: %v got: %v", index, tc.expected, result)
	}

}
