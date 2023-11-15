package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists))
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}

	if l > r {
		return nil
	}

	mid := (l + r) >> 1
	return mergeTwoLists(merge(lists, l, mid), merge(lists, mid+1, r))
}

func mergeTwoLists(a, b *ListNode) *ListNode {
	if a == nil || b == nil {
		if a != nil {
			return a
		}
		return b
	}

	head := &ListNode{Val: 0}
	tail, aPtr, bPtr := head, a, b

	// 循环对比大小，小数连接到tail
	for aPtr != nil && bPtr != nil {
		if aPtr.Val < bPtr.Val {
			tail.Next = aPtr
			aPtr = aPtr.Next
		} else {
			tail.Next = bPtr
			bPtr = bPtr.Next
		}

		tail = tail.Next
	}

	if aPtr != nil {
		tail.Next = aPtr
	} else {
		tail.Next = bPtr
	}

	return head.Next
}
