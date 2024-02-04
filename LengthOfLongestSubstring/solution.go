package LengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	prevSeenChars := make(map[rune]int)
	leftPointer := 0
	length := 0

	for currentIndex, char := range s {
		if positionInMap, seen := prevSeenChars[char]; seen && leftPointer <= positionInMap {
			leftPointer = positionInMap + 1
		}
		prevSeenChars[char] = currentIndex
		length = max(length, currentIndex-leftPointer+1)
	}

	return length
}
