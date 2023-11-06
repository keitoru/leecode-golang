package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre int

func isValidBST(root *TreeNode) bool {
	pre = math.MinInt
	return bts(root)

}

func bts(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !bts(root.Left) {
		return false
	}

	if root.Val <= pre {
		return false
	}

	pre = root.Val

	return bts(root.Right)
}
