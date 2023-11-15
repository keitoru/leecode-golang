package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	length := m + n
	left, right := -1, -1
	aStart, bStart := 0, 0

	for i := 0; i <= length/2; i++ {
		left = right

		if aStart < m && (bStart >= n || nums1[aStart] < nums2[bStart]) {
			right = nums1[aStart]
			aStart++
		} else {
			right = nums2[bStart]
			bStart++
		}
	}

	if length&1 == 0 {
		return float64(left+right) / 2.0
	}

	return float64(right)

}
