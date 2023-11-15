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

	pq:=
		var compare	func(a,b int) int
	compare = func(a, b int) int {
		return
	}

	for k, v := range m {

	}
}
