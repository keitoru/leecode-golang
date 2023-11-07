package main

func subsets(nums []int) (ans [][]int) {
	var res [][]int

	var dfs func(i int, list []int)
	dfs = func(i int, list []int) {

		tmp := make([]int, len(list))
		copy(tmp, list)

		res = append(res, tmp)

		for j := i; j < len(nums); j++ {
			list = append(list, nums[j])
			dfs(j+1, list)
			list = list[:len(list)-1]
		}
	}
	dfs(0, []int{})
	return
}
