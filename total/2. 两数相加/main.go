package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry, sum := 0, 0

	head := &ListNode{}
	cur := head

	p, q := l1, l2

	for p != nil || q != nil || carry > 0 {
		var v1, v2 int

		if p == nil {
			v1 = 0
		} else {
			v1 = p.Val
			p = p.Next
		}

		if q == nil {
			v2 = 0
		} else {
			v2 = q.Val
			q = q.Next
		}

		sum = (v1 + v2 + carry) % 10
		carry = (v1 + v2 + carry) / 10

		next := &ListNode{
			Val: sum,
		}

		cur.Next = next
		cur = cur.Next

	}

	return head.Next

}
