package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canJump(nums []int) bool {
	maxLen := 0

	for i := 0; i < len(nums); i++ {
		if i > maxLen {
			return false
		}

		maxLen = max(maxLen, i+nums[i])
	}

	return true
}
