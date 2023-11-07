package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := dfs(root.Left)
		right := dfs(root.Right)

		innerMaxSum := left + right + root.Val
		maxSum = max(maxSum, innerMaxSum)

		//outputSum才是递归函数的
		//outputSum只选取了左子树或者右子树
		//因此它的值可以传给它的父节点组成更长的合法的路径
		//如果左右子树sum都小于0，那就只计算当前节点的值
		outputMaxSum := root.Val + max(max(0, left), right)

		return max(outputMaxSum, 0)
	}

	dfs(root)

	return maxSum
}
