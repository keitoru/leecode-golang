package main

import (
	"fmt"
	"math"
)

func firstMissingPositive(nums []int) int {
	N := len(nums)
	for i := 0; i < N; i++ {
		if nums[i] <= 0 {
			nums[i] = N + 1
		}
	}

	for i := 0; i < N; i++ {
		num := int(math.Abs(float64(nums[i])))
		if num <= N && num >= 1 {
			nums[num-1] = -1 * int(math.Abs(float64(nums[num-1])))
		}
	}

	for i := 0; i < N; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return N + 1

}

func main() {
	nums := []int{1, 1}
	res := firstMissingPositive(nums)
	fmt.Println(res)
}
