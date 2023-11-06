package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getLength(head *ListNode) int {
	n := 0
	cur := head
	for ; cur != nil; cur = cur.Next {
		n++
	}

	return n
}

//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	length := getLength(head)
//	dummy := &ListNode{0, head}
//	cur := dummy
//	for i := 0; i < length-n; i++ {
//		cur = cur.Next
//	}
//
//	cur.Next = cur.Next.Next
//	return dummy.Next
//}

// 快慢指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast := head
	slow := dummy

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}

func main() {

}
