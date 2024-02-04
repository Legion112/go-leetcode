package RemoveDuplicatesFromSortedArray

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name         string
		nums         *[]int
		expectedNums *[]int // Only the first expectedLen elements will be checked.
	}{
		{
			name:         "Example 1",
			nums:         &[]int{1, 1, 2},
			expectedNums: &[]int{1, 2},
		},
		{
			name:         "Example 2",
			nums:         &[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedNums: &[]int{0, 1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLen := removeDuplicates(*tt.nums)
			if gotLen != len(*tt.expectedNums) {
				t.Errorf("%s: expected length %d, got %d", tt.name, len(*tt.expectedNums), gotLen)
			}
			for i := 0; i < len(*tt.expectedNums); i++ {
				if (*tt.nums)[i] != (*tt.expectedNums)[i] {
					t.Errorf("%s: at index %d, expected num %d, got %d", tt.name, i, (*tt.expectedNums)[i], (*tt.nums)[i])
					break
				}
			}
		})
	}
}
