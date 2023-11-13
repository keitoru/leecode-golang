package main

import "fmt"

func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	i, j, k := n-2, n-1, n-1

	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}

		nums[i], nums[k] = nums[k], nums[i]
	}

	for i, j := j, n-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {
	n := []int{1, 2, 3, 5, 7, 6, 8, 4}
	nextPermutation(n)
	fmt.Println(n)
}
