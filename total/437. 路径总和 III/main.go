package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var prefixMap map[int]int
var target int

func pathSum(root *TreeNode, targetSum int) int {
	prefixMap = make(map[int]int, 100)
	target = targetSum

	prefixMap[0] = 1
	return recur(root, 0)
}

func recur(node *TreeNode, curSum int) int {
	if node == nil {
		return 0
	}

	res := 0
	curSum += node.Val

	res += prefixMap[curSum-target]
	prefixMap[curSum] = prefixMap[curSum] + 1

	left := recur(node.Left, curSum)
	right := recur(node.Right, curSum)

	res = res + left + right

	prefixMap[curSum] = prefixMap[curSum] - 1
	return res
}
