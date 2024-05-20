package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pseudoHead := ListNode{}
	pseudoHead.Next = &ListNode{}
	curr := pseudoHead.Next
	a, b, sum, carry := 0, 0, 0, 0
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			a = l1.Val
		}
		if l2 != nil {
			b = l2.Val
		}
		sum = a + b + carry
		curr.Val = sum % 10
		carry = sum / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if l1 == nil && l2 == nil && carry == 0 {
			curr.Next = nil
		} else {
			curr.Next = &ListNode{}
		}
		curr = curr.Next
		a, b = 0, 0
	}
	return pseudoHead.Next
}
