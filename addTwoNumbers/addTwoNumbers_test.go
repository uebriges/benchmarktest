package addTwoNumbers

import "testing"

func BenchmarkAddition(b *testing.B) {
	l1 := ListNode{}
	l2 := ListNode{}

	l1.Val = 2
	l1.Next = &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}
	l2.Val = 5
	l2.Next = &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}

	addTwoNumbers(&l1, &l2)
}
