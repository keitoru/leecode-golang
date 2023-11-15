package main

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

// 链表翻转
func reverse(head *ListNode) *ListNode {
	pre, curr := new(ListNode), head
	for curr != nil {
		next := curr.Next //next 指向下一个节点,保存当前节点后面的链表
		curr.Next = pre   //将当前节点next域指向前一个节点   null<-1<-2<-3<-4
		pre = curr        //pre 指针向后移动。pre指向当前节点
		curr = next       //curr指针向后移动。下一个节点变成当前节点
	}
	return pre
}

// 链表分区为已翻转部分+待翻转部分+未翻转部分
func reverseKGroup(head *ListNode, k int) *ListNode {
	//定义一个假的节点
	dummy := &ListNode{Val: 0}
	dummy.Next = head

	//初始化pre和end都指向dummy。
	//pre指每次要翻转的链表的头结点的上一个节点。end指每次要翻转的链表的尾节点
	pre, end := dummy, dummy

	for end.Next != nil {
		//循环k次，找到需要翻转的链表的结尾,这里每次循环要判断end是否等于空,因为如果为空，end.next会报空指针异常。
		//dummy->1->2->3->4->5 若k为2，循环2次，end指向2
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}

		//如果end==null，即需要翻转的链表的节点数小于k，不执行翻转
		if end == nil {
			break
		}

		// 准备翻转
		// 先记录下end.next,方便后面链接链表
		start, next := pre.Next, end.Next
		// 断开链表
		end.Next = nil
		// 翻转[start,end]区间
		pre.Next = reverse(start)
		//翻转后头节点变到最后。通过.next把断开的链表重新链接。
		start.Next = next
		//将pre换成下次要翻转的链表的头结点的上一个节点。即start
		pre = start
		//将pre换成下次要翻转的链表的头结点的上一个节点。即start
		end = pre
	}

	return dummy.Next
}
