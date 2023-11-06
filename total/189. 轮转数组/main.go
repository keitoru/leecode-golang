package main

import "fmt"

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}

func rotate(nums []int, k int) {
	index := k % len(nums)
	reverse(nums)
	reverse(nums[:index])
	reverse(nums[index:])
}

//func rotate(nums []int, k int) {
//	index := len(nums) - k%len(nums)
//	tmp := nums[index:]
//	copy(nums, append(tmp, nums[:index]...))
//}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 10)
	fmt.Println(nums)
}
