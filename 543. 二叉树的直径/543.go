package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxLen int

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxLen = 0

	depth(root)

	return maxLen
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := depth(root.Left)
	r := depth(root.Right)

	maxLen = max(maxLen, l+r)

	return max(l, r) + 1
}
