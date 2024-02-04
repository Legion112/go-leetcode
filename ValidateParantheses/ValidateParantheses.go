package ValidateParantheses

import "leetcode/DataStructures"

func isValid(s string) bool {
	var stack DataStructures.Stack[rune]
	var oposit = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	var last rune
	for _, parentheses := range s {
		openParentheses, isOpen := oposit[parentheses]
		if !isOpen {
			stack.Push(parentheses)
		} else {
			last, _ = stack.Pop()
			if last != openParentheses {
				return false
			}
		}
	}
	return len(stack) == 0
}
