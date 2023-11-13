package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func partitionLabels(s string) []int {
	maxPos := map[byte]int{}
	for i := 0; i < len(s); i++ {
		maxPos[s[i]] = i
	}

	res := []int{}
	start := 0
	scannedCharMaxPos := 0
	for i := 0; i < len(s); i++ {
		curCharMaxPos := maxPos[s[i]]

		scannedCharMaxPos = max(curCharMaxPos, scannedCharMaxPos)
		if i == scannedCharMaxPos {
			res = append(res, i-start+1)
			start = i + 1
		}
	}

	return res
}
