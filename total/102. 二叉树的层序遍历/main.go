package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0, 100)
	queue := []*TreeNode{}

	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		var level []int

		N := len(queue)
		for i := 0; i < N; i++ {
			node := queue[i]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[N:]

		res = append(res, level)
	}
	return res
}
