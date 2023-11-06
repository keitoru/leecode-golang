package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	s1, s2 := make(map[int]bool), make(map[int]bool)
	for _, i1 := range nums1 {
		s1[i1] = true
	}
	for _, i2 := range nums2 {
		s2[i2] = true
	}

	var a, b []int
	for v := range s1 {
		if !s2[v] {
			a = append(a, v)
		}
	}
	for v := range s2 {
		if !s1[v] {
			b = append(b, v)
		}
	}

	return [][]int{a, b}
}
