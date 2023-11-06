package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归解法
//func swapPairs(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	next := head.Next
//
//	head.Next = swapPairs(next.Next)
//	next.Next = head
//
//	return next
//}

// 非递归解法
func swapPairs(head *ListNode) *ListNode {
	pre := &ListNode{
		Val:  0,
		Next: head,
	}

	tmp := pre
	for tmp.Next != nil && tmp.Next.Next != nil {
		start, end := tmp.Next, tmp.Next.Next
		tmp.Next = end
		start.Next = end.Next
		end.Next = start

		tmp = start
	}

	return pre.Next
}

func main() {

}
