package MergeTwoSortedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummyHead ListNode
	currentHead := &dummyHead

	for l1 != nil || l2 != nil {
		if l1 == nil && l2 != nil {
			currentHead.Next = l2
			return dummyHead.Next
		}
		if l1 != nil && l2 == nil {
			currentHead.Next = l1
			return dummyHead.Next
		}

		if l1.Val < l2.Val {
			currentHead.Next = l1
			l1 = l1.Next
		} else {
			currentHead.Next = l2
			l2 = l2.Next
		}
		currentHead = currentHead.Next
	}

	return dummyHead.Next
}
