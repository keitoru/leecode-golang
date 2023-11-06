package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

// dfs 按照 「根结点 -> 右子树 -> 左子树」 的顺序访问
func rightSideView(root *TreeNode) []int {
	res = []int{}
	dfs(root, 0)
	return res

}

func dfs(root *TreeNode, depth int) {
	if root == nil {
		return
	}

	if depth == len(res) {
		res = append(res, root.Val)
	}

	depth++

	dfs(root.Right, depth)
	dfs(root.Left, depth)
}
