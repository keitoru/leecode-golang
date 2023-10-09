package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pA := headA
	pB := headB

	for pA != pB {
		if pA != nil {
			pA = pA.Next
		} else {
			pA = headB
		}

		if pB != nil {
			pB = pB.Next
		} else {
			pB = headA
		}
	}

	return pA
}

func main() {
	node := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}

	listA := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  1,
			Next: node,
		},
	}
	listB := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  1,
				Next: node,
			},
		},
	}

	res := getIntersectionNode(listA, listB)
	fmt.Println(res.Val)
}
