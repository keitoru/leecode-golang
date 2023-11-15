package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	res := [][]int{}

	if n < 3 {
		return res
	}

	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		L, R := i+1, n-1
		for L < R {
			if nums[i]+nums[L]+nums[R] == 0 {
				res = append(res, []int{nums[i], nums[L], nums[R]})

				// 相同的值就跳过
				for L < R && nums[L] == nums[L+1] {
					L++
				}

				for L < R && nums[R] == nums[R-1] {
					R--
				}

				L++
				R--
			} else if nums[i]+nums[L]+nums[R] > 0 {
				R--
			} else {
				L++
			}

		}
	}

	return res
}

/*func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	res := [][]int{}

	for k := 0; k < len(nums)-1; k++ {
		if nums[k] > 0 {
			break
		}

		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		i, j := k+1, len(nums)-1

		for i < j {
			sum := nums[k] + nums[i] + nums[j]

			if sum < 0 {
				i++
				for i-1 < j && nums[i-1] == nums[i] {
					i++
				}

				//fmt.Println(i)
			} else if sum > 0 {
				j--
				for i < j+1 && nums[j+1] == nums[j] {
					j--
				}
			} else {
				res = append(res, []int{nums[k], nums[i], nums[j]})

				i++
				for i-1 < j && nums[i-1] == nums[i] {
					i++
				}

				j--
				for i < j+1 && nums[j+1] == nums[j] {
					j--
				}
			}
		}
	}

	return res
}*/

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println(res)
}
