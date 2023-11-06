package main

import "fmt"

func moveZeroes(nums []int) {
	zero := 0

	p, q := 0, 0
	N := len(nums)

	for q < N {
		if nums[q] != zero {
			nums[p], nums[q] = nums[q], nums[p]
			p++
		}

		q++
	}
}

func main() {
	nums := []int{1, 0, 1}
	moveZeroes(nums)
	fmt.Println(nums)
}
