package FindTheIndexOfTheFirstOccurrenceInString

func strStr(haystack string, needle string) int {
	needleCharacters := []rune(needle)
	lastIndex := len(needleCharacters) - 1
	haystackCharacters := []rune(haystack)
	for i := 0; i < len(haystackCharacters) && (i+len(needleCharacters)) <= len(haystackCharacters); i++ {
		index := 0
		for j, needleChar := range needleCharacters {
			index = i + j
			if haystackCharacters[index] != needleChar {
				break
			}
			if j == lastIndex {
				return i
			}
		}
	}
	return -1
}
