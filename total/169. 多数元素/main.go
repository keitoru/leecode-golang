package main

//func majorityElement(nums []int) int {
//	sort.Slice(nums, func(i, j int) bool {
//		return nums[i] < nums[j]
//	})
//
//	return nums[len(nums)/2]
//}

func majorityElement(nums []int) int {
	m := make(map[int]int, 100)

	for _, num := range nums {

		if _, ok := m[num]; !ok {
			m[num] = 0
		}

		m[num]++
	}

	for num, n := range m {
		if n > len(nums)/2 {
			return num
		}
	}

	return -1
}
