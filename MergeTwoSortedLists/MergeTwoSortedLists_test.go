package MergeTwoSortedLists

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name   string
		list1  *ListNode
		list2  *ListNode
		expect *ListNode
	}{
		{
			name:   "Example 1",
			list1:  &ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			list2:  &ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			expect: &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		},
		{
			name:   "Example 2",
			list1:  nil,
			list2:  nil,
			expect: nil,
		},
		{
			name:   "Example 3",
			list1:  nil,
			list2:  &ListNode{0, nil},
			expect: &ListNode{0, nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeTwoLists(tt.list1, tt.list2)
			require.Equal(t, tt.expect, result)
		})
	}
}
