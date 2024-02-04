package DataStructures

type Stack[T any] []T

func (stack *Stack[T]) Push(value T) {
	*stack = append(*stack, value)
}

func (stack *Stack[T]) Pop() (T, bool) {
	if len(*stack) == 0 {
		var value T // init default value of T
		return value, false
	}
	topIndex := len(*stack) - 1
	element := (*stack)[topIndex]
	*stack = (*stack)[:topIndex] //remove top element by slicing it off
	return element, true
}
