package main

import "fmt"

var list []int
var res [][]int

func permute(nums []int) [][]int {
	list = make([]int, len(nums))
	res = make([][]int, 0)
	list = nums
	dfs(0)
	return res
}

func dfs(x int) {
	if x == len(list)-1 {
		tmp := make([]int, len(list))
		copy(tmp, list)
		res = append(res, tmp)
		return
	}

	for i := x; i < len(list); i++ {
		swap(i, x)
		dfs(x + 1)
		swap(i, x)
	}
}

func swap(a, b int) {
	list[a], list[b] = list[b], list[a]
}

func main() {
	nums := []int{1, 2, 3}
	res := permute(nums)
	fmt.Println(res)
}
