package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	N := len(candidates)

	var dfs func(target int, index int, cob []int)
	dfs = func(target int, index int, cob []int) {
		if index == N {
			return
		}

		//fmt.Println("cob:", cob, ",target:", target)

		if target == 0 {
			tmp := make([]int, len(cob))
			copy(tmp, cob)

			res = append(res, tmp)
			return
		}

		for i := index; i < N; i++ {
			if target-candidates[i] >= 0 {
				cob = append(cob, candidates[i])

				dfs(target-candidates[i], i, cob)

				cob = cob[:len(cob)-1]
			}

		}
	}

	dfs(target, 0, []int{})

	return res
}

/*func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	N := len(candidates)
	comb := []int{}

	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == N {
			return
		}

		if target == 0 {
			res = append(res, append([]int{}, comb...))
			return
		}

		dfs(target, idx+1)

		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}

	dfs(target, 0)

	return res
}*/

func main() {
	c := []int{2, 3, 6, 7}
	target := 7

	res := combinationSum(c, target)
	fmt.Println(res)
}
