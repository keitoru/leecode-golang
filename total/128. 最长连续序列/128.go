package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	maxLen := 0
	hashMap := make(map[int]bool)

	for _, num := range nums {
		hashMap[num] = true
	}

	for num := range hashMap {
		if !hashMap[num-1] {
			curNum := num
			curLen := 1

			for hashMap[curNum+1] {
				curLen++
				curNum += 1
			}

			if curLen > maxLen {
				maxLen = curLen
			}
		}
	}

	return maxLen

}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}

	res := longestConsecutive(nums)

	fmt.Println(res)
}
