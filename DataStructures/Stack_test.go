package DataStructures

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStackEverythingPushedComeOutPop(t *testing.T) {
	var stack Stack[rune]

	stack.Push('(')
	element, exist := stack.Pop()
	require.True(t, exist)
	require.Equal(t, '(', element)
}

func TestStackEmptyStackShouldReturnDefaultValueAndFalse(t *testing.T) {
	var stack Stack[rune]

	element, exist := stack.Pop()
	require.False(t, exist)
	var defaultValue rune
	require.Equal(t, defaultValue, element)
}
