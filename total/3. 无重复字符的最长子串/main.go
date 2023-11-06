package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]bool, len(s))
	left := 0
	curLen, maxLen := 0, 0

	for _, ch := range []byte(s) {
		curLen += 1
		for m[ch] {
			delete(m, s[left])
			left += 1
			curLen -= 1
		}

		if curLen > maxLen {
			maxLen = curLen
		}

		m[ch] = true
	}

	return maxLen
}

func main() {
	res := lengthOfLongestSubstring("pwwkew")
	fmt.Println(res)
}
