package TwoSum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		list     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for index, tc := range testCases {
		var result = twoSum(tc.list, tc.target)
		require.ElementsMatchf(
			t,
			tc.expected,
			result,
			"Test index: %d, list: %v, expected: %v, result: %v",
			index,
			tc.list,
			tc.expected,
			result,
		)
	}
}
