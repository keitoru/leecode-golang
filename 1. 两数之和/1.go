package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numsLen := len(nums)
	m := make(map[int]int, numsLen)

	for i, num := range nums {
		subNum := target - num
		if subIndex, ok := m[subNum]; ok {
			return []int{subIndex, i}
		}
		m[num] = i
	}

	return []int{}
}

func main() {
	nums := []int{2, 5, 4, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
}
