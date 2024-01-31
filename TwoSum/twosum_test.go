package TwoSum

import (
	"fmt"
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
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v target: %v expected: %v", tc.list, tc.target, tc.expected), func(t *testing.T) {
			var result = twoSum(tc.list, tc.target)
			require.ElementsMatchf(
				t,
				tc.expected,
				result,
				"List: %v, expected: %v, result: %v",
				tc.list,
				tc.expected,
				result,
			)
		})
	}
}
