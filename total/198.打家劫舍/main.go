package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 子问题：
	// f(k) = 偷 [0..k) 房间中的最大金额

	// f(0) = 0
	// f(1) = nums[0]
	// f(k) = max{ rob(k-1), nums[k-1] + rob(k-2) }

	N := len(nums)

	// 优化前
	//dp := make([]int, N+1)
	//dp[0] = 0
	//dp[1] = nums[0]
	//for k := 2; k <= N; k++ {
	//	dp[k] = max(dp[k-1], nums[k-1]+dp[k-2])
	//}

	// 优化空间
	dp1 := 0
	dp2 := nums[0]
	for k := 2; k <= N; k++ {
		dp1, dp2 = dp2, max(dp2, nums[k-1]+dp1)
	}

	return dp2
}

func main() {
	nums := []int{1, 9, 3, 1}
	res := rob(nums)
	fmt.Println(res)
}
