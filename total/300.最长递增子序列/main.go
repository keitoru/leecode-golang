package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS(nums []int) int {
	N := len(nums)
	dp := make([]int, N)

	res := 0

	for i := 0; i < N; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}

	return res
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	res := lengthOfLIS(nums)
	fmt.Println(res)
}
