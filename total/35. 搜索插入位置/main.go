package main

func searchInsert(nums []int, target int) int {
	pos := 0
	for i, num := range nums {
		if num >= target {
			return i
		}

		pos = i
	}

	return pos + 1
}
