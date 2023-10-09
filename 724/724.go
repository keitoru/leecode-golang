package main

import "fmt"

func pivotIndex(nums []int) int {
	left, right := 0, 0
	movePtr := 0

	for _, num := range nums {
		right += num
	}

	for movePtr < len(nums) {
		right -= nums[movePtr]
		if left == right {
			return movePtr
		}
		left += nums[movePtr]
		movePtr++
	}

	return -1
}

func main() {
	nums := []int{-1, -1, -1, -1, -1, -1}
	i := pivotIndex(nums)
	fmt.Println(i)
}
