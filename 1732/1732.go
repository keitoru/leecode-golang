package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestAltitude(gain []int) int {
	//gain = [-5,1,5,0,-7]
	var maxHeight, curHeight int
	curHeight = 0

	for _, v := range gain {
		curHeight += v
		maxHeight = max(maxHeight, curHeight)
	}

	return maxHeight
}

func main() {
	gain := []int{-5, 1, 5, 0, -7}
	height := largestAltitude(gain)
	fmt.Println(height)
}
