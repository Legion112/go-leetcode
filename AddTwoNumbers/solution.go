package AddTwoNumbers

type ListNode struct {
	Val  uint8
	Next *ListNode
}

func addTwoNumbers(a *ListNode, b *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	carry := dummy.Val

	for a != nil || b != nil || carry > 0 {
		if carry == 0 {
			// Either 'a' or 'b' is not nil, but not both
			if a != nil && b == nil {
				curr.Next = a
				return dummy.Next
			} else if a == nil && b != nil {
				curr.Next = b
				return dummy.Next
			}
		}
		sum := carry

		if a != nil {
			sum += a.Val
			a = a.Next
		}

		if b != nil {
			sum += b.Val
			b = b.Next
		}

		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
	}

	return dummy.Next
}
