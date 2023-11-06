package main

import "fmt"

func subarraySum(nums []int, k int) int {
	numsSum := 0
	prefixNums := map[int]int{0: 1}
	count := 0

	for _, num := range nums {
		// 前缀和
		numsSum += num

		if prefixNums[numsSum-k] > 0 {
			count += prefixNums[numsSum-k]
		}

		prefixNums[numsSum]++

	}

	return count
}

func main() {
	nums := []int{1, 3, 1, 2}
	k := 3

	res := subarraySum(nums, k)
	fmt.Println(res)
}
