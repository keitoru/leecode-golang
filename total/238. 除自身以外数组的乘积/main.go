package main

import "fmt"

func productExceptSelf(nums []int) []int {
	N := len(nums)
	left, right := make([]int, N), make([]int, N)

	left[0] = 1
	for i := 1; i < N; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right[N-1] = 1
	for i := N - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := 0; i < N; i++ {
		left[i] *= right[i]
	}

	return left
}

func main() {
	nums := []int{1, 2, 3, 4}
	res := productExceptSelf(nums)
	fmt.Println(res)

}
