package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxArea(height []int) int {
	N := len(height)
	left, right := 0, N-1

	area := 0

	for left < right {
		if height[left] < height[right] {
			area = max(area, height[left]*(right-left))
			left++
		} else {
			area = max(area, height[right]*(right-left))
			right--
		}
	}

	return area
}

func main() {
	h := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(h)
	fmt.Println(res)
}
