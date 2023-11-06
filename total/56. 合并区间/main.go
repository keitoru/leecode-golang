package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {

	var merged [][]int
	if len(intervals) == 0 {
		return merged
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals); i++ {
		L := intervals[i][0]
		R := intervals[i][1]

		if len(merged) == 0 || merged[len(merged)-1][1] < L {
			merged = append(merged, []int{L, R})
		} else {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], R)
		}
	}

	return merged
}

func main() {
	i := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	res := merge(i)
	fmt.Println(res)
}
