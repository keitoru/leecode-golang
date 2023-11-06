package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	tmp := 0
	maxNum := nums[0]

	for _, num := range nums {
		tmp = max(num, tmp+num)
		maxNum = max(maxNum, tmp)
	}

	return maxNum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := maxSubArray(nums)
	fmt.Println(res)
}
