package main

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func canPartition(nums []int) bool {
	n := len(nums)

	if n < 2 {
		return false
	}

	// 计算整个数组的sum 以及最大元素maxNum
	sum, maxNum := 0, 0
	for _, num := range nums {
		sum += num
		maxNum = max(maxNum, num)
	}

	//如果 sum 是奇数
	//则不可能将数组分割成元素和相等的两个子集
	if sum%2 != 0 {
		return false
	}

	// 需要判断是否可以从数组中选出一些数字，使得这些数字的和等于target
	target := sum / 2
	// 如果maxNum>target，剩余元素之和一定小于target
	if maxNum > target {
		return false
	}

	dp := make([]bool, target+1)
	dp[0] = true

	for i := 0; i < n; i++ {
		num := nums[i]
		for j := target; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}

	return dp[target]
}
