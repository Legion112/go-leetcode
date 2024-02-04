package RemoveDuplicatesFromSortedArray

func removeDuplicates(nums []int) int {
	leftIndex := 0
	leftValue := (nums)[0]
	swaps := 0
	for _, rightValue := range (nums)[1:] {
		if leftValue != rightValue {
			leftIndex++
			(nums)[leftIndex] = rightValue
			leftValue = rightValue
			swaps++
		}
	}
	return swaps + 1
}
