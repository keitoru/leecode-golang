package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	// cut递归终止条件：只有一个节点了，直接返回此节点
	if head == nil || head.Next == nil {
		return head
	}

	fast, slow := head.Next, head

	// 用快慢指针找到中点slow
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	tmp := slow.Next
	slow.Next = nil // 将链表截断

	left := sortList(head)
	right := sortList(tmp)

	// 合并
	h := &ListNode{Val: 0}
	res := h

	for left != nil && right != nil {
		if left.Val < right.Val {
			h.Next = left
			left = left.Next
		} else {
			h.Next = right
			right = right.Next
		}

		h = h.Next
	}

	if left != nil {
		h.Next = left
	} else {
		h.Next = right
	}

	return res.Next

}

func main() {

}
