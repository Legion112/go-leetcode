package AddTwoNumbers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	testCases := map[string]struct {
		first    *ListNode
		second   *ListNode
		expected *ListNode
	}{
		"342+465=708": {
			first:    &ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			second:   &ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			expected: &ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
		"0+0=0": {
			first:    &ListNode{0, nil},
			second:   &ListNode{0, nil},
			expected: &ListNode{0, nil},
		},
		"9999999+9999=89990001": {
			first:    &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}},
			second:   &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}},
			expected: &ListNode{8, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}},
		},
	}

	for name, tc := range testCases {
		t.Run(fmt.Sprintf(name), func(t *testing.T) {
			result := addTwoNumbers(tc.first, tc.second)
			assert.Equal(t, tc.expected, result, "Structs should be equal expected: %v got: %v", tc.expected, result)
		})
	}

}
