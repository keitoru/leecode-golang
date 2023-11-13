package main

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, num := range nums {
		if _, ok := m[num]; ok {
			m[num]++
		} else {
			m[num] = 1
		}
	}

	for k, v := range m {

	}
}
