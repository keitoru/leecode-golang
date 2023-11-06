package main

import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProduct(nums []int) int {
	N := len(nums)

	maxNum, imax, imin := math.MinInt, 1, 1

	for i := 0; i < N; i++ {
		if nums[i] < 0 {
			imax, imin = imin, imax
		}
		imax = max(imax*nums[i], nums[i])
		imin = min(imin*nums[i], nums[i])

		maxNum = max(maxNum, imax)
	}

	return maxNum
}

func main() {

}
