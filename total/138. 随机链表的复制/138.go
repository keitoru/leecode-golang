package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 拼接 + 拆分
func copyRandomList(head *Node) *Node {
	if head == nil {
		return &Node{}
	}

	//1. 复制各节点，并构建拼接链表
	cur := head
	for cur != nil {
		tmp := &Node{
			Val:  cur.Val,
			Next: cur.Next,
		}
		cur.Next = tmp

		cur = tmp.Next
	}

	//2. 构建各新节点的 random 指向
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}

		cur = cur.Next.Next
	}

	// 3. 拆分两链表
	cur = head.Next
	res := head.Next
	pre := head
	for cur.Next != nil {
		pre.Next = cur.Next
		cur.Next = cur.Next.Next

		cur = cur.Next
		pre = pre.Next
	}

	pre.Next = nil

	return res
}

// hash
//func copyRandomList(head *Node) *Node {
//	hashMap := make(map[*Node]*Node)
//
//	cur := head
//
//	for cur != nil {
//		hashMap[cur] = &Node{Val: cur.Val}
//		cur = cur.Next
//	}
//
//	cur = head
//	for cur != nil {
//
//		if cur.Next != nil {
//			hashMap[cur].Next = hashMap[cur.Next]
//		}
//
//		if cur.Random != nil {
//			hashMap[cur].Random = hashMap[cur.Random]
//		}
//		cur = cur.Next
//	}
//
//	return hashMap[head]
//}
