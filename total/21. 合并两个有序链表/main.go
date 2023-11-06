package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p, q := list1, list2
	head := &ListNode{}
	cur := head

	for p != nil || q != nil {
		if p == nil {
			cur.Next = q
			break
		}

		if q == nil {
			cur.Next = p
			break
		}

		if p.Val <= q.Val {
			cur.Next = p
			p = p.Next
		} else {
			cur.Next = q
			q = q.Next
		}

		cur = cur.Next
	}

	return head.Next
}
