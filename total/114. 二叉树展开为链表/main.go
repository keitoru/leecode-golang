package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func flatten(root *TreeNode) {
//	for root != nil {
//		if root.Left != nil {
//			pre := root.Left
//
//			for pre.Right != nil {
//				pre = pre.Right
//			}
//
//			pre.Right = root.Right
//			root.Right = root.Left
//			root.Left = nil
//
//		}
//
//		root = root.Right
//	}
//}

func flatten(root *TreeNode) {
	toVisit := []*TreeNode{}
	cur, pre := root, new(TreeNode)
	pre = nil

	for cur != nil || len(toVisit) > 0 {
		for cur != nil {
			toVisit = append(toVisit, cur)
			cur = cur.Right
		}

		cur = toVisit[len(toVisit)-1]
		if cur.Left == nil || cur.Left == pre {
			toVisit = toVisit[:len(toVisit)-1]

			cur.Right = pre
			cur.Left = nil

			pre = cur
			cur = nil
		} else {
			cur = cur.Left
		}
	}
}
