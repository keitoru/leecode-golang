package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return dfs(nums, 0, len(nums)-1)
}

func dfs(nums []int, low, high int) *TreeNode {
	if low > high {
		return nil
	}

	mid := (high + low) / 2

	root := &TreeNode{
		Val: nums[mid],
	}
	root.Left = dfs(nums, low, mid-1)
	root.Right = dfs(nums, mid+1, high)

	return root
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	res := sortedArrayToBST(nums)

	for res.Right != nil {
		fmt.Println(res.Val)

		res = res.Right
	}
}
